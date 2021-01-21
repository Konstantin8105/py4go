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
class itemList ( dict ):

  def add ( self, ID, item ):
    
    if ID in self:
      raise RuntimeError( 'ID ' + str(ID) + ' already exists in ' + type(self).__name__ )

    self[ID] = item

  def get ( self, IDs ):

    if isinstance(IDs,int):
      return self[IDs]
    elif isinstance(IDs,list):
      return [ self[ID] for ID in IDs ]
      
    raise RuntimeError('illegal argument for itemList.get')  

  def getIndices ( self, IDs ):
    
    if isinstance(IDs,int):
      return list(self.keys()).index( IDs )
    elif isinstance(IDs,list):
      return [ list(self.keys()).index( ID ) for ID in IDs ]
      
    raise RuntimeError('illegal argument for itemList.getIndices')  
