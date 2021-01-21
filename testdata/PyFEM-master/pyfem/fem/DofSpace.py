############################################################################
#  This Python file is part of PyFEM, the code that accompanies the book:  #
#                                                                          #
#    'Non-Linear Finite Element Analysis of Solids and Structures'         #
#    R. de Borst, M.A. Crisfield, J.J.C. Remmers and C.V. Verhoosel        #
#    John Wiley and Sons, 2012, ISBN 978-0470666449                        #
#                                                                          #
#  The code is written by J.J.C. Remmers, C.V. Verhoosel and R. de Borst.  #
#                                                                          #
#  The latest stable version can be downloaded from the web-site:          #
#     http://www.wiley.com/go/deborst                                      #
#                                                                          #
#  A github repository, with the most up to date version of the code,      #
#  can be found here:                                                      #
#     https://github.com/jjcremmers/PyFEM                                  #
#                                                                          #
#  The code is open source and intended for educational and scientific     #
#  purposes only. If you use PyFEM in your research, the developers would  #
#  be grateful if you could cite the book.                                 #  
#                                                                          #
#  Disclaimer:                                                             #
#  The authors reserve all rights but do not guarantee that the code is    #
#  free from errors. Furthermore, the authors shall not be liable in any   #
#  event caused by the use of the program.                                 #
############################################################################

from numpy import array, dot, zeros
import scipy.linalg

from scipy.sparse.linalg   import spsolve
from scipy.sparse.linalg   import eigsh
from pyfem.util.itemList   import itemList
from pyfem.util.fileParser import readNodeTable
from pyfem.util.logger     import getLogger
from pyfem.fem.Constrainer import Constrainer

from copy import deepcopy

logger = getLogger()


class DofSpace:

  def __init__ ( self, elements ):

    self.dofTypes = elements.getDofTypes()
    self.dofs     = array( list(range( len(elements.nodes) * len(self.dofTypes))) ).reshape( ( len(elements.nodes), len(self.dofTypes) ) )
    self.nodes    = elements.nodes
    
    #Create the ID map
    self.IDmap = itemList()
    for ind,ID in enumerate(elements.nodes):
      self.IDmap.add( ID, ind )

    self.allConstrainedDofs = []

#
#
#

  def __str__ ( self ):
    return str(self.dofs)

#
#
#

  def __len__ ( self ):
    return len(self.dofs.flatten())

#
#
#

  def setConstrainFactor( self , fac , loadCase = "All_" ):

    if loadCase == "All_":
      for name in self.cons.constrainedFac.keys():
        self.cons.constrainedFac[name] = fac
    else:
      self.cons.constrainedFac[loadCase] = fac

#
#
#
    
  def readFromFile( self, fname ):
      
    logger.info("Reading constraints ..........")

    nodeTable = readNodeTable( fname , "NodeConstraints" , self.nodes )
   
    self.cons = self.createConstrainer( nodeTable )
              
  def createConstrainer ( self, nodeTables ):
        
    cons = Constrainer(len(self))
    
    for nodeTable in nodeTables:
      
      label = nodeTable.subLabel

      cons.constrainedDofs[label] = []
      cons.constrainedVals[label] = []
      cons.constrainedFac [label] = 1.0
      
      for item in nodeTable.data:
        nodeID  = item[1]
        dofType = item[0]
        val     = item[2]

        if not nodeID in self.nodes:
          raise RuntimeError('Node ID ' + str(nodeID) + ' does not exist')

        ind = self.IDmap.get( nodeID )

        if dofType not in self.dofTypes:
          raise RuntimeError('DOF type "' + dofType + '" does not exist')
      
        if len(item) == 3:          
          dofID = self.dofs[ind,self.dofTypes.index(dofType)]
        
          cons.addConstraint(dofID,val,label)
        else:
          slaveNodeID  = item[4]
          slaveDofType = item[3]
          factor       = item[5]

          if not slaveNodeID in self.nodes:
            raise RuntimeError('Node ID ' + str(slaveNodeID) + ' does not exist')

          slaveInd = self.IDmap.get( slaveNodeID )

          if slaveDofType not in self.dofTypes:
            raise RuntimeError('DOF type "' + slaveDofType + '" does not exist')
      
          slaveDof = self.dofs[slaveInd,self.dofTypes.index(slaveDofType)]
      
          dofID = self.dofs[ind,self.dofTypes.index(dofType)]
                  
          cons.addConstraint(dofID , [ val , slaveDof , factor ] , label )
                    
    cons.flush()
        
    return cons

#-------------------------------------------------------------------------------
#
#-------------------------------------------------------------------------------

  def getForType ( self, nodeIDs, dofType ):
  
    '''Returns all dofIDs for given dofType for a list of nodes'''
   
    return self.dofs[self.IDmap.get( nodeIDs ), self.dofTypes.index(dofType)]
      
#-------------------------------------------------------------------------------
#
#-------------------------------------------------------------------------------

  def getForTypes( self, nodeIDs, dofTypes ):
  
    '''Returns all dofIDs for given list of dofType for a list of nodes'''

    dofs = []
        
    for node in nodeIDs:
      for dofType in dofTypes:
        dofs.append(self.dofs[self.IDmap.get( node ),self.dofTypes.index(dofType)])
      
    return dofs

#-------------------------------------------------------------------------------
#
#-------------------------------------------------------------------------------

  def get ( self, nodeIDs ):
  
    '''Returns all dofIDs for a list of nodes'''
    
    return self.dofs[self.IDmap.get(nodeIDs)].flatten()
    
#-------------------------------------------------------------------------------
#
#-------------------------------------------------------------------------------

  def copyConstrainer( self , dofTypes: list = None ):
  
    newCons = deepcopy(self.cons)
       
    if type(dofTypes) is str:
      dofTypes = [dofTypes]
          
    for dofType in dofTypes:
      for iDof in self.dofs[:,self.dofTypes.index(dofType)]:
        for label in newCons.constrainedFac.keys():
          newCons.addConstraint(iDof,0.0,label)
                  
    newCons.flush()
                  
    return newCons

#-------------------------------------------------------------------------------
#  
#-------------------------------------------------------------------------------

  def solve ( self, A, b, constrainer = None ):

    '''Solves the system Ax = b using the internal constraints matrix.
       Returns the total solution vector x.'''
    
    if constrainer is None:
      constrainer = self.cons
      
    if len(A.shape) == 2:

      a = zeros(len(self))
      
      constrainer.addConstrainedValues( a )

      A_constrained = constrainer.C.transpose() * (A * constrainer.C )
      b_constrained = constrainer.C.transpose() * ( b - A * a )
            
      x_constrained = spsolve( A_constrained, b_constrained )

      x = constrainer.C * x_constrained
      
      constrainer.addConstrainedValues( x )
          
    elif len(A.shape) == 1:
      x = b / A

      constrainer.setConstrainedValues( x )
   
    return x
    
    
#-------------------------------------------------------------------------------
#
#-------------------------------------------------------------------------------

  def eigensolve( self, A , B , count=5 ):

    '''Calculates the first count eigenvlaues and eigenvectors of a
       system with ( A lambda B ) x '''
       
    A_constrained = dot( dot( self.cons.C.transpose(), A ), self.cons.C )
    B_constrained = dot( dot( self.cons.C.transpose(), B ), self.cons.C )

    eigvals , eigvecs = eigsh( A_constrained, count , B_constrained , sigma = 0. , which = 'LM' )

    x = zeros(shape=(len(self),count))

    for i,psi in enumerate(eigvecs.transpose()):
      x[:,i] = self.cons.C * psi
      
    return eigvals,x

#-------------------------------------------------------------------------------
#
#-------------------------------------------------------------------------------

  def norm ( self, r ):
    
    return scipy.linalg.norm( self.cons.C.transpose() * r )
