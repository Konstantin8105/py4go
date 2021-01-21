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

from .Element import Element
from pyfem.util.transformations import getRotationMatrix
from pyfem.util.shapeFunctions  import getElemShapeData	
from pyfem.elements.Composite   import Laminate,stressTransformation

from numpy import zeros, ones, dot, array, eye, outer, mat, empty,sqrt
from scipy.linalg import norm
from math import atan2, sin, cos, tan

#==============================================================================
#
#==============================================================================

class postProcessPoint:
  pass

#------------------------------------------------------------------------------
#
#------------------------------------------------------------------------------

class Plate ( Element ):

  #dofs per element
  dofTypes = [ 'u' , 'v' , 'w' , 'rx' , 'ry' ]

#------------------------------------------------------------------------------
#
#------------------------------------------------------------------------------

  def __init__ ( self, elnodes , props ):
    Element.__init__( self, elnodes , props )

    self.material = Laminate( props )

    self.A = self.material.getA()
    self.B = self.material.getB()
    self.D = self.material.getD()

    Ashear = self.material.getAshear()

    self.A44 = Ashear[0,0]
    self.A45 = Ashear[0,1]
    self.A55 = Ashear[1,1]

    self.inertia = self.material.getMassInertia()

    self.initPostProcessing()

    self.outputLabels = self.postProcess[0].labels+self.postProcess[1].labels

#------------------------------------------------------------------------------
#
#------------------------------------------------------------------------------

  def __type__ ( self ):
    return name

#------------------------------------------------------------------------------
#
#------------------------------------------------------------------------------

  def getTangentStiffness ( self, elemdat ):

    sData = getElemShapeData( elemdat.coords )
    
    for d in sData:

      stiff = zeros(shape=(20,20))

      N11 = self.A[0,0]*d.dhdx[:,0] + self.A[0,2]*d.dhdx[:,1]
      N21 = self.A[0,1]*d.dhdx[:,1] + self.A[0,2]*d.dhdx[:,0]
      N41 = self.B[0,0]*d.dhdx[:,0] + self.B[0,2]*d.dhdx[:,1]
      N51 = self.B[0,1]*d.dhdx[:,1] + self.B[0,2]*d.dhdx[:,0]

      N12 = self.A[0,1]*d.dhdx[:,0] + self.A[1,2]*d.dhdx[:,1]
      N22 = self.A[1,1]*d.dhdx[:,1] + self.A[1,2]*d.dhdx[:,0]
      N42 = self.B[0,1]*d.dhdx[:,0] + self.B[1,2]*d.dhdx[:,1]
      N52 = self.B[1,1]*d.dhdx[:,1] + self.B[1,2]*d.dhdx[:,0]

      N16 = self.A[0,2]*d.dhdx[:,0] + self.A[2,2]*d.dhdx[:,1]
      N26 = self.A[1,2]*d.dhdx[:,1] + self.A[2,2]*d.dhdx[:,0]
      N46 = self.B[0,2]*d.dhdx[:,0] + self.B[2,2]*d.dhdx[:,1]
      N56 = self.B[1,2]*d.dhdx[:,1] + self.B[2,2]*d.dhdx[:,0]

      M11 = self.B[0,0]*d.dhdx[:,0] + self.B[0,2]*d.dhdx[:,1]
      M21 = self.B[0,1]*d.dhdx[:,1] + self.B[0,2]*d.dhdx[:,0]
      M41 = self.D[0,0]*d.dhdx[:,0] + self.D[0,2]*d.dhdx[:,1]
      M51 = self.D[0,1]*d.dhdx[:,1] + self.D[0,2]*d.dhdx[:,0]

      M12 = self.B[0,1]*d.dhdx[:,0] + self.B[1,2]*d.dhdx[:,1]
      M22 = self.B[1,1]*d.dhdx[:,1] + self.B[1,2]*d.dhdx[:,0]
      M42 = self.D[0,1]*d.dhdx[:,0] + self.D[1,2]*d.dhdx[:,1]
      M52 = self.D[1,1]*d.dhdx[:,1] + self.D[1,2]*d.dhdx[:,0]

      M16 = self.B[0,2]*d.dhdx[:,0] + self.B[2,2]*d.dhdx[:,1]
      M26 = self.B[1,2]*d.dhdx[:,1] + self.B[2,2]*d.dhdx[:,0]
      M46 = self.D[0,2]*d.dhdx[:,0] + self.D[2,2]*d.dhdx[:,1]
      M56 = self.D[1,2]*d.dhdx[:,1] + self.D[2,2]*d.dhdx[:,0]

      for i in range(4):
        for j in range(4):
 
          #K1
          stiff[5*i+0,5*j+0] += d.dhdx[i,0]*N11[j]+d.dhdx[i,1]*N16[j]
          stiff[5*i+1,5*j+0] += d.dhdx[i,0]*N16[j]+d.dhdx[i,1]*N12[j]

          stiff[5*i+3,5*j+0] += d.dhdx[i,0]*M11[j]+d.dhdx[i,1]*M16[j]
          stiff[5*i+4,5*j+0] += d.dhdx[i,0]*M16[j]+d.dhdx[i,1]*M12[j]
          
          #K2
          stiff[5*i+0,5*j+1] += d.dhdx[i,0]*N21[j]+d.dhdx[i,1]*N26[j]
          stiff[5*i+1,5*j+1] += d.dhdx[i,0]*N26[j]+d.dhdx[i,1]*N22[j]

          stiff[5*i+3,5*j+1] += d.dhdx[i,0]*M21[j]+d.dhdx[i,1]*M26[j]
          stiff[5*i+4,5*j+1] += d.dhdx[i,0]*M26[j]+d.dhdx[i,1]*M22[j]

          #K4
          stiff[5*i+0,5*j+3] += d.dhdx[i,0]*N41[j]+d.dhdx[i,1]*N46[j]
          stiff[5*i+1,5*j+3] += d.dhdx[i,0]*N46[j]+d.dhdx[i,1]*N42[j]
          stiff[5*i+3,5*j+3] += d.dhdx[i,0]*M41[j]+d.dhdx[i,1]*M46[j]
          stiff[5*i+4,5*j+3] += d.dhdx[i,0]*M46[j]+d.dhdx[i,1]*M42[j]

          #K5
          stiff[5*i+0,5*j+4] += d.dhdx[i,0]*N51[j]+d.dhdx[i,1]*N56[j]
          stiff[5*i+1,5*j+4] += d.dhdx[i,0]*N56[j]+d.dhdx[i,1]*N52[j]
          stiff[5*i+3,5*j+4] += d.dhdx[i,0]*M51[j]+d.dhdx[i,1]*M56[j]
          stiff[5*i+4,5*j+4] += d.dhdx[i,0]*M56[j]+d.dhdx[i,1]*M52[j]

      elemdat.stiff += stiff * d.weight

    sData = getElemShapeData( elemdat.coords , -1 )
    
    for d in sData:

      stiff = zeros(shape=(20,20))

      Q41 = self.A55*d.h
      Q42 = self.A45*d.h
      Q51 = self.A45*d.h
      Q52 = self.A44*d.h
      Q31 = self.A55*d.dhdx[:,0]+self.A45*d.dhdx[:,1]
      Q32 = self.A45*d.dhdx[:,0]+self.A44*d.dhdx[:,1]

      for i in range(4):
        for j in range(4):
 
          #K3
          stiff[5*i+2,5*j+2] += d.dhdx[i,0]*Q31[j]+d.dhdx[i,1]*Q32[j]
          stiff[5*i+3,5*j+2] += d.h[i]*Q31[j]
          stiff[5*i+4,5*j+2] += d.h[i]*Q32[j]

          #K4
          stiff[5*i+2,5*j+3] += d.dhdx[i,0]*Q41[j]+d.dhdx[i,1]*Q42[j]
          stiff[5*i+3,5*j+3] += d.h[i]*Q41[j]
          stiff[5*i+4,5*j+3] += d.h[i]*Q42[j]

          #K5
          stiff[5*i+2,5*j+4] += d.dhdx[i,0]*Q51[j]+d.dhdx[i,1]*Q52[j]
          stiff[5*i+3,5*j+4] += d.h[i]*Q51[j]
          stiff[5*i+4,5*j+4] += d.h[i]*Q52[j]

      elemdat.stiff += stiff * d.weight
     
#------------------------------------------------------------------------------
#
#------------------------------------------------------------------------------

  def getInternalForce ( self, elemdat ):

    sData = getElemShapeData( elemdat.coords , -1 )

    elemdat.outlabel.append(self.outputLabels)
    elemdat.outdata  = zeros( shape=(len(elemdat.nodes),6) )

    eps0  = zeros(3)
    kappa = zeros(3)

    for iData in sData:
      eps0[0] = dot(iData.dhdx[:,0],elemdat.state[0:20:5])
      eps0[1] = dot(iData.dhdx[:,1],elemdat.state[1:20:5])
      eps0[2] = dot(iData.dhdx[:,1],elemdat.state[0:20:5])+\
                dot(iData.dhdx[:,0],elemdat.state[1:20:5])

      kappa[0]= dot(iData.dhdx[:,0],elemdat.state[3:20:5])
      kappa[1]= dot(iData.dhdx[:,1],elemdat.state[4:20:5])
      kappa[2]= dot(iData.dhdx[:,1],elemdat.state[3:20:5])+\
                dot(iData.dhdx[:,0],elemdat.state[4:20:5])

      for i,pp in enumerate(self.postProcess):
        eps   = eps0 + pp.z*kappa
        sigma = stressTransformation( dot(pp.Qbar,eps) , pp.theta )

        self.appendNodalOutput( pp.labels , sigma )

#------------------------------------------------------------------------------
#
#------------------------------------------------------------------------------

  def getMassMatrix ( self, elemdat ):
      
    sData = getElemShapeData( elemdat.coords )

    mass = zeros( shape=(20,20) )

    for d in sData:
      for i in range(4):
        for j in range(4):
          shp = d.h[i]*d.h[j]
          mass[5*i+0,5*j+0] += self.inertia[0]*shp
          mass[5*i+0,5*j+3] += self.inertia[1]*shp
          mass[5*i+1,5*j+1] += self.inertia[0]*shp
          mass[5*i+1,5*j+4] += self.inertia[1]*shp
          mass[5*i+2,5*j+2] += self.inertia[0]*shp
          mass[5*i+3,5*j+0] += self.inertia[1]*shp
          mass[5*i+3,5*j+3] += self.inertia[0]*shp
          mass[5*i+3,5*j+4] += self.inertia[1]*shp
          mass[5*i+4,5*j+1] += self.inertia[1]*shp
          mass[5*i+4,5*j+3] += self.inertia[1]*shp
          mass[5*i+4,5*j+4] += self.inertia[0]*shp

      elemdat.mass += mass * d.weight

#------------------------------------------------------------------------------
#
#------------------------------------------------------------------------------

  def initPostProcessing( self ):
 
    layerCount = self.material.layerCount()

    self.postProcess = []
  
    pp        = postProcessPoint()
    pp.z      = -0.5*self.material.thick
    pp.Qbar   = self.material.getQbar(0)
    pp.theta  = self.material.layers[0].theta
    pp.labels = ["s11bot","s22bot","s12bot"]
    self.postProcess.append(pp)

    pp        = postProcessPoint()
    pp.z      = 0.5*self.material.thick
    pp.Qbar   = self.material.getQbar(layerCount-1)
    pp.theta  = self.material.layers[layerCount-1].theta
    pp.labels = ["s11top","s22top","s12top"]
    self.postProcess.append(pp)

    
