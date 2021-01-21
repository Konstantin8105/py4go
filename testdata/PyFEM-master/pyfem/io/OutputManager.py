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
class OutputManager:

  def __init__( self , props , globdat ):

    self.outman = []

    outputModules = props.outputModules

    for name in outputModules:
   
      props.currentModule = name

      ioType = name

      if hasattr( props , name):
       moduleProps = getattr( props, name )
       if hasattr( moduleProps , "type" ):
         ioType = moduleProps.type

      exec("from pyfem.io."+ioType+" import "+ioType)

      self.outman.append(eval(ioType+"( props , globdat )"))

  def run( self , props , globdat ):

    for i,output in enumerate(self.outman):
      output.run( props , globdat )
