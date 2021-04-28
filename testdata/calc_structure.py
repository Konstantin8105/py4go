import math
from scipy.special import gammaln
from scipy.stats import gamma as gammadist
import numpy as np

import ANYstructure.SN_curve_parameters as snc

class Structure():
    '''
    Setting the properties for the plate and the stiffener. Takes a dictionary as argument.
    '''
    def __init__(self, main_dict, *args, **kwargs):
        super(Structure,self).__init__()
        self.main_dict = main_dict
        self.plate_th = main_dict['plate_thk'][0]
        self.web_height = main_dict['stf_web_height'][0]
        self.web_th = main_dict['stf_web_thk'][0]
        self.flange_width = main_dict['stf_flange_width'][0]
        self.flange_th = main_dict['stf_flange_thk'][0]
        self.mat_yield = main_dict['mat_yield'][0]
        self.span = main_dict['span'][0]
        self.spacing = main_dict['spacing'][0]
        self.structure_type = main_dict['structure_type'][0]
        self.sigma_y1=main_dict['sigma_y1'][0]
        self.sigma_y2=main_dict['sigma_y2'][0]
        self.sigma_x=main_dict['sigma_x'][0]
        self.tauxy=main_dict['tau_xy'][0]
        self.plate_kpp = main_dict['plate_kpp'][0]
        self.stf_kps = main_dict['stf_kps'][0]
        self.km1 = main_dict['stf_km1'][0]
        self.km2 = main_dict['stf_km2'][0]
        self.km3 = main_dict['stf_km3'][0]
        self.stiffener_type=main_dict['stf_type'][0]

        self.sigma_y = self.sigma_y2 + (self.sigma_y1-self.sigma_y2)\
                                       *(min(0.25*self.span,0.5*self.spacing)/self.span)
        try:
            self.girder_lg=main_dict['girder_lg'][0]
        except KeyError:
            self.girder_lg = 10
        try:
            self.pressure_side = main_dict['press_side'][0]
        except KeyError:
            self.pressure_side = 'p'

    def __str__(self):
        '''
        Returning all properties.
        '''
        return \
            str(
            '\n Plate field span:              ' + str(round(self.span,1)) + ' meters' +
            '\n Stiffener spacing:             ' + str(self.spacing*1000)+' mm'+
            '\n Plate thickness:               ' + str(self.plate_th*1000)+' mm'+
            '\n Stiffener web height:          ' + str(self.web_height*1000)+' mm'+
            '\n Stiffener web thickness:       ' + str(self.web_th*1000)+' mm'+
            '\n Stiffener flange width:        ' + str(self.flange_width*1000)+' mm'+
            '\n Stiffener flange thickness:    ' + str(self.flange_th*1000)+' mm'+
            '\n Material yield:                ' + str(self.mat_yield/1e6)+' MPa'+
            '\n Structure type/stiffener type: ' + str(self.structure_type)+'/'+(self.stiffener_type)+
            '\n Plate fixation paramter,kpp:   ' + str(self.plate_kpp) + ' ' +
            '\n Stf. fixation paramter,kps:    ' + str(self.stf_kps) + ' ' +
            '\n Global stress, sig_y1/sig_y2:  ' + str(round(self.sigma_y1,1))+'/'+str(round(self.sigma_y2,1))+ ' MPa' +
            '\n Global stress, sig_x:          ' + str(round(self.sigma_x,1)) + ' MPa' +
            '\n Global shear, tau_xy:          ' + str(round(self.tauxy,1)) + ' MPa' +
            '\n km1,km2,km3:                   ' + str(self.km1)+'/'+str(self.km2)+'/'+str(self.km3)+
            '\n Pressure side (p-plate/s-stf): ' + str(self.pressure_side) + ' ')

    def get_one_line_string(self):
        ''' Returning a one line string. '''
        return 'pl_'+str(round(self.spacing*1000, 1))+'x'+str(round(self.plate_th*1000,1))+' stf_'+self.stiffener_type+\
               str(round(self.web_height*1000,1))+'x'+str(round(self.web_th*1000,1))+'+'\
               +str(round(self.flange_width*1000,1))+'x'+\
               str(round(self.flange_th*1000,1))

    def get_report_stresses(self):
        'Return the stresses to the report'
        return 'sigma_y1: '+str(round(self.sigma_y1,1))+' sigma_y2: '+str(round(self.sigma_y2,1))+ \
               ' sigma_x: ' + str(round(self.sigma_x,1))+' tauxy: '+ str(round(self.tauxy,1))

    def get_extended_string(self):
        ''' Some more information returned. '''
        return 'span: '+str(round(self.span,4))+' structure type: '+ self.structure_type + ' stf. type: ' + \
               self.stiffener_type + ' pressure side: ' + self.pressure_side
    
    def get_sigma_y1(self):
        '''
        Return sigma_y1
        :return:
        '''
        return self.sigma_y1
    def get_sigma_y2(self):
        '''
        Return sigma_y2
        :return:
        '''
        return self.sigma_y1
    def get_sigma_x(self):
        '''
        Return sigma_x
        :return:
        '''
        return self.sigma_y1
    def get_tau_xy(self):
        '''
        Return tau_xy
        :return:
        '''
        return self.tauxy
    def get_s(self):
        '''
        Return the spacing
        :return:
        '''
        return self.spacing
    def get_pl_thk(self):
        '''
        Return the plate thickness
        :return:
        '''
        return self.plate_th
    def get_web_h(self):
        '''
        Return the web heigh
        :return:
        '''
        return self.web_height
    def get_web_thk(self):
        '''
        Return the spacing
        :return:
        '''
        return self.web_th
    def get_fl_w(self):
        '''
        Return the flange width
        :return:
        '''
        return self.flange_width
    def get_fl_thk(self):
        '''
        Return the flange thickness
        :return:
        '''
        return self.flange_th
    def get_fy(self):
        '''
        Return material yield
        :return:
        '''
        return self.mat_yield
    def get_span(self):
        '''
        Return the span
        :return:
        '''
        return self.span
    def get_lg(self):
        '''
        Return the girder length
        :return:
        '''
        return self.girder_lg
    def get_kpp(self):
        '''
        Return var
        :return:
        '''
        return self.plate_kpp
    def get_kps(self):
        '''
        Return var
        :return:
        '''
        return self.stf_kps
    def get_km1(self):
        '''
        Return var
        :return:
        '''
        return self.km1
    def get_km2(self):
        '''
        Return var
        :return:
        '''
        return self.km2
    def get_km3(self):
        '''
        Return var
        :return:
        '''
        return self.km3
    def get_side(self):
        '''
        Return the checked pressure side.
        :return: 
        '''
        return self.pressure_side
    def get_tuple(self):
        ''' Return a tuple of the plate stiffener'''
        return (self.spacing, self.plate_th, self.web_height, self.web_th, self.flange_width,
                self.flange_th, self.span, self.girder_lg, self.stiffener_type)

    def get_section_modulus(self, efficient_se = None, dnv_table = False):
        '''
        Returns the section modulus.
        :param efficient_se: 
        :return: 
        '''
        #Plate. When using DNV table, default values are used for the plate
        b1 = self.spacing if efficient_se==None else efficient_se
        tf1 = self.plate_th

        #Stiffener
        tf2 = self.flange_th
        b2 = self.flange_width
        h = self.flange_th+self.web_height+self.plate_th
        tw = self.web_th
        hw = self.web_height

        # cross section area
        Ax = tf1 * b1 + tf2 * b2 + hw * tw

        assert Ax != 0, 'Ax cannot be 0'
        # distance to center of gravity in z-direction
        ez = (tf1 * b1 * tf1 / 2 + hw * tw * (tf1 + hw / 2) + tf2 * b2 * (tf1 + hw + tf2 / 2)) / Ax

        #ez = (tf1 * b1 * (h - tf1 / 2) + hw * tw * (tf2 + hw / 2) + tf2 * b2 * (tf2 / 2)) / Ax
        # moment of inertia in y-direction (c is centroid)

        Iyc = (1 / 12) * (b1 * math.pow(tf1, 3) + b2 * math.pow(tf2, 3) + tw * math.pow(hw, 3))
        Iy = Iyc + (tf1 * b1 * math.pow(tf1 / 2, 2) + tw * hw * math.pow(tf1+hw / 2, 2) +
             tf2 * b2 * math.pow(tf1+hw+tf2 / 2, 2)) - Ax * math.pow(ez, 2)

        # elastic section moduluses y-axis
        Wey1 = Iy / (h - ez)
        Wey2 = Iy / ez
        return Wey1, Wey2
    def get_plasic_section_modulus(self):
        '''
        Returns the plastic section modulus
        :return:
        '''
        tf1 = self.plate_th
        tf2 = self.flange_th
        b1 = self.spacing
        b2 = self.flange_width
        h = self.flange_th+self.web_height+self.plate_th
        tw = self.web_th
        hw = self.web_height

        Ax = tf1 * b1 + tf2 * b2 + (h-tf1-tf2) * tw

        ezpl = (Ax/2-b1*tf1)/tw+tf1

        az1 = h-ezpl-tf1
        az2 = ezpl-tf2

        Wy1 = b1*tf1*(az1+tf1/2) + (tw/2)*math.pow(az1,2)
        Wy2 = b2*tf2*(az2+tf2/2)+(tw/2)*math.pow(az2,2)

        return Wy1+Wy2
    def get_shear_center(self):
        '''
        Returning the shear center
        :return:
        '''
        tf1 = self.plate_th
        tf2 = self.flange_th
        b1 = self.spacing
        b2 = self.flange_width
        h = self.flange_th+self.web_height+self.plate_th
        tw = self.web_th
        hw = self.web_height
        Ax = tf1 * b1 + tf2 * b2 + (h-tf1-tf2) * tw
        # distance to center of gravity in z-direction
        ez = (b2*tf2*tf2/2 + tw*hw*(tf2+hw/2)+tf1*b1*(tf2+hw+tf1/2)) / Ax

        # Shear center:
        # moment of inertia, z-axis
        Iz1 = tf1 * math.pow(b1, 3)
        Iz2 = tf2 * math.pow(b2, 3)
        ht = h - tf1 / 2 - tf2 / 2
        return (Iz1 * ht) / (Iz1 + Iz2) + tf2 / 2 - ez
    def get_moment_of_intertia(self, efficent_se=None):
        '''
        Returning moment of intertia.
        :return:
        '''
        tf1 = self.plate_th
        b1 = self.spacing if efficent_se==None else efficent_se
        h = self.flange_th+self.web_height+self.plate_th
        tw = self.web_th
        hw = self.web_height
        tf2 = self.flange_th
        b2 = self.flange_width

        Ax = tf1 * b1 + tf2 * b2 + (h-tf1-tf2) * tw
        Iyc = (1 / 12) * (b1 * math.pow(tf1, 3) + b2 * math.pow(tf2, 3) + tw * math.pow(hw, 3))
        ez = (tf1 * b1 * (h - tf1 / 2) + hw * tw * (tf2 + hw / 2) + tf2 * b2 * (tf2 / 2)) / Ax
        Iy = Iyc + (tf1 * b1 * math.pow(tf2 + hw + tf1 / 2, 2) + tw * hw * math.pow(tf2 + hw / 2, 2) +
             tf2 * b2 * math.pow(tf2 / 2, 2)) - Ax * math.pow(ez, 2)
        return Iy

    def get_structure_prop(self):
        return self.main_dict

    def get_structure_type(self):
        return self.structure_type

    def get_stiffener_type(self):
        return self.stiffener_type

    def get_shear_area(self):
        '''
        Returning the shear area in [m^2]
        :return:
        '''
        return ((self.flange_th*self.web_th) + (self.web_th*self.plate_th) + (self.web_height*self.web_th))

    def set_main_properties(self, main_dict):
        '''
        Resettting all properties
        :param input_dictionary:
        :return:
        '''

        self.main_dict = main_dict
        self.plate_th = main_dict['plate_thk'][0]
        self.web_height = main_dict['stf_web_height'][0]
        self.web_th = main_dict['stf_web_thk'][0]
        self.flange_width = main_dict['stf_flange_width'][0]
        self.flange_th = main_dict['stf_flange_thk'][0]
        self.mat_yield = main_dict['mat_yield'][0]
        self.span = main_dict['span'][0]
        self.spacing = main_dict['spacing'][0]
        self.structure_type = main_dict['structure_type'][0]
        self.sigma_y1=main_dict['sigma_y1'][0]
        self.sigma_y2=main_dict['sigma_y2'][0]
        self.sigma_x=main_dict['sigma_x'][0]
        self.tauxy=main_dict['tau_xy'][0]
        self.plate_kpp = main_dict['plate_kpp'][0]
        self.stf_kps = main_dict['stf_kps'][0]
        self.km1 = main_dict['stf_km1'][0]
        self.km2 = main_dict['stf_km2'][0]
        self.km3 = main_dict['stf_km3'][0]
        self.stiffener_type=main_dict['stf_type'][0]
        try:
            self.girder_lg=main_dict['girder_lg'][0]
        except KeyError:
            self.girder_lg = 10
        try:
            self.pressure_side = main_dict['press_side'][0]
        except KeyError:
            self.pressure_side = 'p'

    def set_stresses(self,sigy1,sigy2,sigx,tauxy):
        '''
        Setting the global stresses.
        :param sigy1:
        :param sigy2:
        :param sigx:
        :param tauxy:
        :return:
        '''
        self.main_dict['sigma_y1'][0]= sigy1
        self.sigma_y1 = sigy1

        self.main_dict['sigma_y2'][0]= sigy2
        self.sigma_y2  = sigy2

        self.main_dict['sigma_x'][0]= sigx
        self.sigma_x = sigx

        self.main_dict['tau_xy'][0]= tauxy
        self.tauxy  = tauxy

    def get_plate_thk(self):
        '''
        Return the plate thickness
        :return:
        '''
        return self.plate_th

    def get_cross_section_area(self, efficient_se = None):
        '''
        Returns the cross section area.
        :return:
        '''
        tf1 = self.plate_th
        tf2 = self.flange_th
        b1 = self.spacing if efficient_se==None else efficient_se
        b2 = self.flange_width
        h = self.flange_th+self.web_height+self.plate_th
        tw = self.web_th
        return tf1 * b1 + tf2 * b2 + (h-tf1-tf2) * tw

    def get_cross_section_centroid_with_effective_plate(self, se):
        '''
        Returns cross section centroid
        :return:
        '''
        # checked with example
        tf1 = self.plate_th
        tf2 = self.flange_th
        b1 = se
        b2 = self.flange_width
        h = self.flange_th+self.web_height+self.plate_th
        tw = self.web_th
        hw = self.web_height
        Ax = tf1 * b1 + tf2 * b2 + hw * tw

        return (tf1 * b1 * tf1/2 + hw * tw * (tf1 + hw / 2) + tf2 * b2 * (tf1+hw+tf2/2)) / Ax

    def get_weight(self):
        '''
        Return the weight.
        :return:
        '''
        return 7850*self.span*(self.spacing*self.plate_th+self.web_height*self.web_th+self.flange_width*self.flange_th)

    def get_weight_width_lg(self):
        '''
        Return the weight including Lg
        :return:
        '''
        pl_area = self.girder_lg*self.plate_th
        stf_area = (self.web_height*self.web_th+self.flange_width*self.flange_th)*(self.girder_lg//self.spacing)
        return (pl_area+stf_area)*7850*self.span

    def set_span(self,span):
        '''
        Setting the span. Used when moving a point.
        :return: 
        '''
        self.span = span
        self.main_dict['span'][0] = span

class CalcScantlings(Structure):
    '''
    This Class does the calculations for the plate fields. 
    Input is a structure object, same as for the structure class.
    The class inherits from Structure class.
    '''
    def __init__(self, main_dict, lat_press = True, category = 'secondary'):
        super(CalcScantlings,self).__init__(main_dict=main_dict)
        self.lat_press = lat_press
        self.category = category
        self._need_recalc = True

    @property
    def need_recalc(self):
        return self._need_recalc

    @need_recalc.setter
    def need_recalc(self, val):
        self._need_recalc = val

    def get_results_for_report(self,lat_press=0):
        '''
        Returns a string for the report.
        :return:
        '''
        buc = [round(res,1) for res in self.calculate_buckling_all(design_lat_press=lat_press)]

        return 'Minimum section modulus:'\
               +str(int(self.get_dnv_min_section_modulus(design_pressure_kpa=lat_press)*1000**3))\
               +'mm^3 '+' Minium plate thickness: '\
               +str(round(self.get_dnv_min_thickness(design_pressure_kpa=lat_press),1))+\
               ' Buckling results: eq7_19: '+str(buc[0])+' eq7_50: '+str(buc[1])+ ' eq7_51: '\
               +str(buc[2])+ ' eq7_52: '+str(buc[3])+ ' eq7_53: '+str(buc[4])

    def calculate_slamming_plate(self, slamming_pressure):
        ''' Slamming pressure input is Pa '''
        ka1 = 1.1
        ka2 = min(max(0.4, self.spacing / self.span), 1)
        ka = math.pow(ka1 - 0.25*ka2,2)
        sigmaf = self.mat_yield/1e6  # MPa
        psl = slamming_pressure/1000  # kPa
        Cd = 1.5

        return 0.0158*ka*self.spacing*1000*math.sqrt(psl/(Cd*sigmaf))

    def calculate_slamming_stiffener(self, slamming_pressure, angle = 90):
        tk = 0
        psl = slamming_pressure / 1000  # kPa
        Pst = psl/2
        sigmaf = self.mat_yield / 1e6  # MPa
        hw, twa, tp, tf, bf, s = [(val - tk) * 1000 for val in [self.web_height, self.web_th, self.plate_th, self.flange_th,
                                                            self.flange_width, self.spacing]]
        ns = 2
        tau_eH = sigmaf/math.sqrt(3)
        h_stf = (self.web_height+self.flange_th)*1000
        f_shr = 0.7
        lbdg = self.span
        lshr = self.span - self.spacing/4000
        dshr = h_stf + tp if 75 <= angle <= 90 else (h_stf + tp)*math.sin(math.radians(angle))
        tw = (f_shr*Pst*s*lshr)/(dshr*tau_eH)

        if self.web_th*1000 < tw:
            return {'tw_req': tw, 'Zp_req':None}
        fpl = 8* (1+(ns/2))
        Zp_req = (1.2*Pst*s*math.pow(lbdg,2)/(fpl*sigmaf)) + \
                  (ns*(1-math.sqrt(1-math.pow(tw/twa,2)))*hw*tw*(hw+tp))/8000

        return {'tw_req': tw, 'Zp_req':Zp_req}

    def check_all_slamming(self, slamming_pressure):
        ''' A summary check of slamming '''

        pl_chk = self.calculate_slamming_plate(slamming_pressure)
        if self.plate_th*1000 < pl_chk:
            chk1 = pl_chk / self.plate_th*1000
            return False, chk1

        stf_res = self.calculate_slamming_stiffener(slamming_pressure)

        if self.web_th*1000 < stf_res['tw_req']:
            chk2 = stf_res['tw_req'] / self.web_th*1000
            return False, chk2

        if stf_res['Zp_req'] is not None:
            eff_pl_sec_mod = self.get_net_effective_plastic_section_modulus()
            if eff_pl_sec_mod < stf_res['Zp_req']:
                chk3 = stf_res['Zp_req']/eff_pl_sec_mod
                return False, chk3

        return True, None

    def get_net_effective_plastic_section_modulus(self, angle = 90):
        ''' Calculated according to Rules for classification: Ships — DNVGL-RU-SHIP Pt.3 Ch.3. Edition July 2017,
            page 83 '''
        tk = 0
        angle_rad = math.radians(angle)
        hw, tw, tp, tf, bf = [(val - tk) * 1000 for val in [self.web_height, self.web_th, self.plate_th, self.flange_th,
                                                            self.flange_width]]
        h_stf = (self.web_height+self.flange_th)*1000
        de_gr = 0
        tw_gr = self.web_th*1000
        hf_ctr = h_stf-0.5*tf if self.get_stiffener_type() != 'L' else h_stf - de_gr - 0.5*tf
        bf_ctr = 0 if self.get_stiffener_type() == 'T' else 0.5*(tf - tw_gr)
        beta = 0.5
        gamma = (1 + math.sqrt(3+12*beta))/4

        Af = 0 if self.get_stiffener_type() == 'FB' else bf*tf

        if 75 <= angle <= 90:
            zpl = (hw*tw*(hw+tp)/2000) + ( (2*gamma-1) * Af * ((hf_ctr + tp/2)) / 1000)
        elif angle < 75:
            zpl = (hw*tw*(hw+tp)/2000)+\
                  ( (2*gamma-1) * Af * ((hf_ctr + tp/2) * math.sin(angle_rad) - bf_ctr*math.cos(angle_rad)) / 1000)

        return zpl

    def get_dnv_min_section_modulus(self, design_pressure_kpa):
        ''' Section modulus according to DNV rules '''

        design_pressure = design_pressure_kpa
        fy = self.mat_yield / 1e6
        fyd = fy/1.15

        sigma_jd = math.sqrt(math.pow(self.sigma_x,2)+math.pow(self.sigma_y,2)-
                             self.sigma_x*self.sigma_y+3*math.pow(self.tauxy,2))

        sigma_pd2 = fyd-sigma_jd  # design_bending_stress_mpa

        kps = self.stf_kps  # 1 is clamped, 0.9 is simply supported.
        km_sides = min(self.km1,self.km3)  # see table 3 in DNVGL-OS-C101 (page 62)
        km_middle = self.km2  # see table 3 in DNVGL-OS-C101 (page 62)

        Zs = ((math.pow(self.span, 2) * self.spacing * design_pressure) /
              (min(km_middle, km_sides) * (sigma_pd2) * kps)) * math.pow(10, 6)

        return max(math.pow(15, 3) / math.pow(1000, 3), Zs / math.pow(1000, 3))

    def get_dnv_min_thickness(self, design_pressure_kpa):
        '''
        Return minimum thickness in mm
        :param design_pressure_kpa:
        :return:
        '''

        design_pressure = design_pressure_kpa
        #print(self.sigma_x)
        sigma_jd = math.sqrt(math.pow(self.sigma_x,2)+math.pow(self.sigma_y,2)-
                             self.sigma_x*self.sigma_y+3*math.pow(self.tauxy,2))
        fy = self.mat_yield / 1000000
        fyd = fy/1.15
        sigma_pd1 = min(1.3*(fyd-sigma_jd), fyd)
        sigma_pd1 = abs(sigma_pd1)
        #print(fyd, sigma_jd, fyd)
        if self.category == 'secondary':
            t0 = 5
        else:
            t0 = 7

        t_min = (14.3 * t0) / math.sqrt(fyd)

        ka = math.pow(1.1 - 0.25  * self.spacing/self.span, 2)

        if ka > 1:
            ka =1
        elif ka<0.72:
            ka = 0.72

        assert sigma_pd1 > 0, 'sigma_pd1 must be negative | current value is: ' + str(sigma_pd1)
        t_min_bend = (15.8 * ka * self.spacing * math.sqrt(design_pressure)) / \
                     math.sqrt(sigma_pd1 *self.plate_kpp)

        if self.lat_press:
            return max(t_min, t_min_bend)
        else:
            return t_min

    def get_minimum_shear_area(self, pressure):
        '''
        Calculating minimum section area according to ch 6.4.4.

        Return [m^2]
        :return:
        '''
        #print('SIGMA_X ', self.sigma_x)
        l = self.span
        s = self.spacing
        fy = self.mat_yield
        fyd = (fy/1.15)/1e6 #yield strength
        sigxd = self.sigma_x #design membrane stresses, x-dir

        taupds = 0.577*math.sqrt(math.pow(fyd, 2) - math.pow(sigxd, 2))

        As = ((l*s*pressure)/(2*taupds)) * math.pow(10,3)

        return As/math.pow(1000,2)

    def is_acceptable_sec_mod(self, section_module, pressure):
        '''
        Checking if the result is accepable.
        :param section_module:
        :param pressure:
        :return:
        '''

        return min(section_module) >= self.get_dnv_min_section_modulus(pressure)

    def is_acceptable_shear_area(self, shear_area, pressure):
        '''
        Returning if the shear area is ok.
        :param shear_area:
        :param pressure:
        :return:
        '''

        return shear_area >= self.get_minimum_shear_area(pressure)

    def get_plate_efficent_b(self,design_lat_press=0,axial_stress=50,
                                 trans_stress_small=100,trans_stress_large=100):
        '''
        Simple buckling calculations according to DNV-RP-C201
        :return:
        '''

        #7.2 Forces in the idealised stiffened plate

        s = self.spacing #ok
        t = self.plate_th #ok
        l = self.span #ok

        E = 2.1e11 #ok

        pSd = design_lat_press*1000
        sigy1Sd =trans_stress_large*1e6
        sigy2Sd =trans_stress_small*1e6
        sigxSd = axial_stress*1e6

        fy = self.mat_yield #ok

        #7.3 Effective plate width
        alphap=0.525*(s/t)*math.sqrt(fy/E) # reduced plate slenderness, checked not calculated with ex
        alphac = 1.1*(s/t)*math.sqrt(fy/E) # checked not calculated with example
        mu6_9 = 0.21*(alphac-0.2)

        if alphac<=0.2: kappa = 1 # eq6.7, all kappa checked not calculated with example
        elif 0.2<alphac<2: kappa = (1/(2*math.pow(alphac,2)))*(1+mu6_9+math.pow(alphac,2)
                                                               -math.sqrt(math.pow(1+mu6_9+math.pow(alphac,2),2)
                                                                          -4*math.pow(alphac,2))) # ok
        else: kappa=(1/(2*math.pow(alphac,2)))+0.07 # ok

        ha = 0.05*(s/t)-0.75
        assert ha>= 0,'ha must be larger than 0'
        kp = 1 if pSd<=2*((t/s)**2)*fy else 1-ha*((pSd/fy)-2*(t/s)**2)

        sigyR=( (1.3*t/l)*math.sqrt(E/fy)+kappa*(1-(1.3*t/l)*math.sqrt(E/fy)))*fy*kp # checked not calculated with example
        l1 = min(0.25*l,0.5*s)

        sig_min, sig_max = min(sigy1Sd,sigy2Sd),max(sigy1Sd,sigy2Sd) # self-made
        sigySd = sig_min+(sig_max-sig_min)*(1-l1/l) # see 6.8, page 15

        ci = 1-s/(120*t) if (s/t)<=120 else 0 # checked not calculated with example

        Cxs = (alphap-0.22)/math.pow(alphap,2) if alphap > 0.673 else 1 # reduction factor longitudinal
        # eq7.16, reduction factor transverse, compression (positive) else tension

        Cys = math.sqrt(1-math.pow(sigySd/sigyR,2) + ci*((sigxSd*sigySd)/(Cxs*fy*sigyR))) if sigySd >= 0 \
            else min(0.5*(math.sqrt(4-3*math.pow(sigySd/fy,2))+sigySd/fy),1) #ok, checked

        #7.7.3 Resistance parameters for stiffeners
        return s * Cxs * Cys # 7.3, eq7.13, che

    def calculate_buckling_all(self,design_lat_press=0.0, checked_side = 'p'):
        '''
        Simple buckling calculations according to DNV-RP-C201
        :return:
        '''
        #7.2 Forces in the idealised stiffened plate
        As = self.web_height*self.web_th+self.flange_width*self.flange_th #checked
        s = self.spacing #ok
        t = self.plate_th #ok
        l = self.span #ok
        tf = self.flange_th
        tw = self.web_th
        hw = self.web_height
        bf = self.flange_width
        fy = self.mat_yield  # ok
        stf_type = self.get_stiffener_type()

        E = 2.1e11 #ok
        Lg = 10 #girder length, ok
        mc = 13.3  # assume continous stiffeners

        pSd = design_lat_press*1000
        tauSd = self.tauxy*1e6
        sigy1Sd =self.sigma_y1*1e6
        sigy2Sd =self.sigma_y2*1e6
        sigxSd = self.sigma_x*1e6


        #7.3 Effective plate width
        alphap=0.525*(s/t)*math.sqrt(fy/E) # reduced plate slenderness, checked not calculated with ex
        alphac = 1.1*(s/t)*math.sqrt(fy/E) # eq 6.8 checked not calculated with example
        mu6_9 = 0.21*(alphac-0.2)

        #kappa chapter 6.3
        if alphac<=0.2: kappa = 1 # eq6.7, all kappa checked not calculated with example
        elif 0.2<alphac<2: kappa = (1/(2*math.pow(alphac,2))) * (1+mu6_9+math.pow(alphac,2)
                                                                 - math.sqrt(math.pow(1+mu6_9+math.pow(alphac,2),2)
                                                                             -4*math.pow(alphac,2))) # ok
        else: kappa=(1/(2*math.pow(alphac,2)))+0.07 # ok
        #end kappa

        ha = 0.05*(s/t)-0.75 #eq 6.11 - checked, ok

        assert ha>= 0,'ha must be larger than 0'
        kp = 1 if pSd<=2*math.pow(t/s,2)*fy else max(1-ha*((pSd/fy)-2*math.pow(t/s,2)),0) #eq 6.10, checked

        sigyR=( (1.3*t/l)*math.sqrt(E/fy)+kappa*(1-(1.3*t/l)*math.sqrt(E/fy)))*fy*kp # eq 6.6 checked

        sigyRd = sigyR / 1.15 #eq 6.5 checked, ok


        # plate resistance check
        ksp = math.sqrt(1-3*math.pow(tauSd/(fy/1),2)) #eq7.20 ch7.4, checked ok

        l1 = min(0.25*l,0.5*s)
        sig_min, sig_max = min(sigy1Sd,sigy2Sd),max(sigy1Sd,sigy2Sd) # self-made
        sigySd = sig_min+(sig_max-sig_min)*(1-l1/l) # see 6.8, page 15

        if not sigySd<=sigyRd:
            return [float('inf'),0,0,0,0]

        try:
            psi = sigy2Sd/sigy1Sd # eq. 7.11 checked, if input is 0, the psi is set to 1
        except ZeroDivisionError:
            psi = 1

        Is = self.get_moment_of_intertia()  # moment of intertia full plate width
        Ip = math.pow(t,3)*s/10.9 # checked not calculated with example

        kc = 2*(1+math.sqrt(1+(10.9*Is)/(math.pow(t,3)*s))) # checked not calculated with example
        kg = 5.34+4*math.pow((l/Lg),2) if l<=Lg else 5.34*math.pow(l/Lg,2)+4 # eq 7.5 checked not calculated with example
        kl = 5.34+4*math.pow((s/l),2) if l>=s else 5.34*math.pow(s/l,2)+4 # eq7.7 checked not calculated with example

        taucrg = kg*0.904*E*math.pow(t/l,2) # 7.2 critical shear stress, checked not calculated with example
        taucrl = kl*0.904*E*math.pow(t/s,2) # 7.2 critical chear stress, checked not calculated with example
        tautf = (tauSd - taucrg) if  tauSd>taucrl/1.15 else 0 # checked not calculated with example

        #7.6 Resistance of stiffened panels to shear stresses (page 20)
        taucrs = (36*E/(s*t*math.pow(l,2)))*((Ip*math.pow(Is,3))**0.25) # checked not calculated with example
        tauRd = min(fy/(math.sqrt(3)*1.15), taucrl/1.15,taucrs/1.15)# checked not calculated with example

        ci = 1-s/(120*t) if (s/t)<=120 else 0 # checked ok
        Cxs = (alphap-0.22)/math.pow(alphap,2) if alphap>0.673 else 1 # reduction factor longitudinal, ok

        # eq7.16, reduction factor transverse, compression (positive) else tension

        Cys = math.sqrt(1-math.pow(sigySd/sigyR,2)+ci*((sigxSd*sigySd)/(Cxs*fy*sigyR))) if sigySd >= 0 \
            else min(0.5*(math.sqrt(4-3*math.pow(sigySd/fy,2))+sigySd/fy),1) #eq 7.16, ok, checked
        #7.7.3 Resistance parameters for stiffeners

        se = s * Cxs * Cys # 7.3, eq7.13, checked
        zp = self.get_cross_section_centroid_with_effective_plate(se) - t / 2  # ch7.5.1 page 19
        zt = (t / 2 + hw + tf) - zp  # ch 7.5.1 page 19

        Ie = self.get_moment_of_intertia(efficent_se=se) #ch7.5.1 effective moment of inertia.
        Wep = Ie/zp #as def in eq7.71
        Wes = Ie/zt #as def in eq7.71

        C0 = (Wes * fy * mc) / (kc * E * math.pow(t, 2) * s)  # 7.2 checked not calculated with example
        p0 = (0.6+0.4*psi)*C0*sigy1Sd if psi>-1.5 else 0 # 7.2 checked not calculated with example

        qSd = (pSd + p0) * s  # checked not calculated with example

        Ae = As+se*t #ch7.7.3 checked, ok

        W = min(Wes,Wep) #eq7.75 text, checked
        pf = (12*W/(math.pow(l,2)*s))*(fy/1.15) #checked, ok

        lk = l*(1-0.5*abs(pSd/pf)) #eq7.74, buckling length, checked

        ie = math.sqrt(Ie/Ae) #ch 7.5.1. checked
        fE = math.pow(math.pi,2)*E*math.pow(ie/lk,2) #e7.24, checked

        sigjSD = math.sqrt(math.pow(sigxSd,2)+math.pow(sigySd,2)-sigxSd*sigySd+3*math.pow(tauSd,2)) # eq 7.38, ok
        fEpx = 3.62*E*math.pow(t/s,2) # eq 7.42, checked, ok
        fEpy = 0.9*E*math.pow(t/s,2) # eq 7.43, checked, ok
        fEpt = 5.0*E*math.pow(t/s,2) # eq 7.44, checked, ok
        c = 2-(s/l) # eq 7.41, checked, ok

        alphae = math.sqrt( (fy/sigjSD) * math.pow(math.pow(sigxSd/fEpx, c)+
                                                   math.pow(sigySd/fEpy, c)+
                                                   math.pow(tauSd/fEpt, c), 1/c)) # eq 7.40, checed, ok.
        fep = fy / math.sqrt(1+math.pow(alphae,4)) # eq 7.39, checked, ok.
        eta = min(sigjSD/fep, 1) # eq. 7.37, chekced

        C = (hw / s) * math.pow(t / tw, 3) * math.sqrt((1 - eta)) # e 7.36, checked ok

        beta = (3*C+0.2)/(C+0.2) # eq 7.35, checked, ok

        Af = self.flange_width*self.flange_th #flange area, ok
        Aw = self.web_height*self.web_th #web area, ok

        ef = 0 if stf_type in ['FB','T'] else self.flange_width/2-self.web_th/2
        #Ipo = (Aw*(ef-0.5*tf)**2/3+Af*ef**2)*10e-4 #polar moment of interia in cm^4
        #It = (((ef-0.5*tf)*tw**3)/3e4)*(1-0.63*(tw/(ef-0.5*tf)))+( (bf*tf)/3e4*(1-0.63*(tf/bf)))/(100**4) #torsonal moment of interia cm^4


        Iz = (1/12)*Af*math.pow(bf,2)+math.pow(ef,2)*(Af/(1+(Af/Aw))) #moment of inertia about z-axis, checked

        G = E/(2*(1+0.3)) #rules for ships Pt.8 Ch.1, page 334
        lT = self.span # Calculated further down
        #print('Aw ',Aw,'Af ', Af,'tf ', tf,'tw ', tw,'G ', G,'E ', E,'Iz ', Iz,'lt ', lT)

        def get_some_data(lT):
            if stf_type in ['T', 'L']:
                fET = beta*(((Aw + Af * math.pow(tf/tw,2)) / (Aw + 3*Af)) * G*math.pow(tw/hw,2))+\
                      (math.pow(math.pi, 2) * E * Iz) / ((Aw/3 + Af)*math.pow(lT,2)) \
                    if bf != 0 \
                    else (beta+2*math.pow(hw/lT,2))*G*math.pow(tw/hw,2) # eq7.32 checked, no example
            else:
                fET = (beta + 2*math.pow(hw/lT,2))*G*math.pow(tw/hw,2) # eq7.34 checked, no example

            alphaT = math.sqrt(fy/fET) #eq7.30. checked

            mu7_29 = 0.35 * (alphaT - 0.6) # eq 7.29. checked

            fr = fy if alphaT<=0.6 else ((1+mu7_29+math.pow(alphaT,2)-math.sqrt( math.pow(1+mu7_29+math.pow(alphaT,2),2)-
                                                                                 4*math.pow(alphaT,2))) /
                                         (2*math.pow(alphaT,2))) * fy
            alpha = math.sqrt(fr / fE) #e7.23, checked.

            mu_tors = 0.35*(alphaT-0.6)
            fT = fy if alphaT <= 0.6 else fy * (1+mu_tors+math.pow(alphaT,2)-math.sqrt(math.pow(1+mu_tors+math.pow(alphaT,2),2)-
                                                                                     4*math.pow(alphaT,2)))/\
                                           (2*math.pow(alphaT,2))

            mu_pl = (0.34 + 0.08 * (zp / ie)) * (alpha - 0.2)
            mu_stf = (0.34 + 0.08 * (zt / ie)) * (alpha - 0.2)
            frp = fy
            frs = fy if alphaT <= 0.6 else fT
            fyp,fys = fy,fy
            #fyps = (fyp*se*t+fys*As)/(se*t+As)
            fks = fr if alpha <= 0.2 else frs * (1+mu_stf+math.pow(alpha,2)-math.sqrt(math.pow(1+mu_stf+math.pow(alpha,2),2)-
                                                                                     4*math.pow(alpha,2)))/\
                                          (2*math.pow(alpha,2))
            #fr = fyps
            fkp = fyp if alpha <= 0.2 else frp * (1+mu_pl+math.pow(alpha,2)-math.sqrt(math.pow(1+mu_pl+math.pow(alpha,2),2)-
                                                                                     4*math.pow(alpha,2)))/\
                                           (2*math.pow(alpha,2))

            return fr, fks, fkp

        u = math.pow(tauSd / tauRd, 2)  # eq7.58. checked.
        fr, fks, fkp = get_some_data(lT=lT*0.4)
        Ms1Rd = Wes*(fr/1.15) #ok, assuming fr calculated with lT=span * 0.4
        NksRd = Ae * (fks / 1.15) #eq7.66, page 22 - fk according to equation 7.26, sec 7.5,
        NkpRd = Ae * (fkp / 1.15)  # checked ok, no ex

        M1Sd = abs((qSd*math.pow(l,2))/12) #ch7.7.1, checked ok

        M2Sd = abs((qSd*l**2)/24) #ch7.7.1, checked ok

        Ne = ((math.pow(math.pi,2))*E*Ae)/(math.pow(lk/ie,2))# eq7.72 , checked ok

        Nrd = Ae * (fy / 1.15) #eq7.65, checked ok

        Nsd = sigxSd * (As + s*t) + tautf * s *t #  Equation 7.1, section 7.2, checked ok

        zstar = 0 #simplification as per 7.7.1 Continuous stiffeners
        MstRd = Wes*(fy/1.15) #eq7.70 checked ok, no ex
        MpRd = Wep*(fy/1.15) #eq7.71 checked ok, no ex

        fr, fks, fkp = get_some_data(lT = lT * 0.8)
        Ms2Rd = Wes*(fr/1.15) #eq7.69 checked ok, no ex
        # print('Nksrd', NksRd, 'Nkprd', NkpRd, 'Ae is', Ae, 'fks is', fks, 'fkp is', fkp,
        #       'alphas are', mu_pl, mu_stf, 'lk', lk, 'lt', lT)

        #print('CENTROID ', 'zp', 'zt', self.get_cross_section_centroid_with_effective_plate(se)*1000,zp,zt)

        eq7_19 = sigySd/(ksp*sigyRd) #checked ok

        # Lateral pressure on plate side:
        if checked_side == 'p':
            # print('eq7_50 = ',Nsd ,'/', NksRd,'+' ,M1Sd,'-' , Nsd ,'*', zstar, '/' ,Ms1Rd,'*',1,'-', Nsd ,'/', Ne,'+', u)
            # print('eq7_51 = ',Nsd,' / ',NkpRd,' - 2 * ',Nsd, '/' ,Nrd,' + ',M1Sd,' - ,',Nsd,' * ',zstar,' / ',MpRd,' * ','1 - ',Nsd,' / ',Ne,' + ',u)
            #print('eq7_52 = ',Nsd,'/', NksRd,'-', 2, '*',Nsd,'/', Nrd,'+',M2Sd,'-', Nsd,'*', zstar,'/',MstRd,'*',1, '-',Nsd,'/', Ne,'+', u)
            max_lfs = []
            ufs = []
            for zstar in np.arange(-zt/2,zp,0.002):
                eq7_50 = (Nsd / NksRd) + (M1Sd - Nsd * zstar) / (Ms1Rd * (1 - Nsd / Ne)) + u
                eq7_51 = (Nsd / NkpRd) - 2 * (Nsd / Nrd) + ((M1Sd - Nsd * zstar) / (MpRd * (1 - (Nsd / Ne)))) + u
                eq7_52 = (Nsd / NksRd) - 2 * (Nsd / Nrd) + ((M2Sd + Nsd * zstar) / (MstRd * (1 - (Nsd / Ne)))) + u
                eq7_53 = (Nsd / NkpRd) + (M2Sd + Nsd * zstar) / (MpRd * (1 - Nsd / Ne))
                max_lfs.append(max(eq7_50, eq7_51, eq7_52, eq7_53))
                ufs.append([eq7_19, eq7_50, eq7_51, eq7_52, eq7_53,zstar])
                #print(zstar, eq7_50, eq7_51, eq7_52, eq7_53, 'MAX LF is: ', max(eq7_50, eq7_51, eq7_52, eq7_53))
            min_of_max_ufs_idx = max_lfs.index(min(max_lfs))
            #print(ufs[min_of_max_ufs_idx])
            return ufs[min_of_max_ufs_idx]
        # Lateral pressure on stiffener side:
        else:
            max_lfs = []
            ufs = []
            for zstar in np.arange(-zt / 2, zp, 0.002):
                eq7_54 = (Nsd / NksRd) - 2 * (Nsd / Nrd) + ((M1Sd + Nsd * zstar) / (MstRd * (1 - (Nsd / Ne)))) + u
                eq7_55 = (Nsd / NkpRd) + ((M1Sd + Nsd * zstar) / (MpRd * (1 - (Nsd / Ne)))) + u
                eq7_56 = (Nsd / NksRd) + ((M2Sd - Nsd * zstar) / (Ms2Rd * (1 - (Nsd / Ne)))) + u
                eq7_57 = (Nsd / NkpRd) - 2 * (Nsd / Nrd) + ((M2Sd - Nsd * zstar) / (MpRd * (1 - (Nsd / Ne)))) + u
                max_lfs.append(max(eq7_54, eq7_55, eq7_56, eq7_57))
                ufs.append([eq7_19, eq7_54, eq7_55, eq7_56, eq7_57, zstar])
                #print('eq7_19, eq7_54, eq7_55, eq7_56, eq7_57')
            min_of_max_ufs_idx = max_lfs.index(min(max_lfs))
            return ufs[min_of_max_ufs_idx]

    def calculate_buckling_plate(self,design_lat_press,axial_stress=20,
                                 trans_stress_small=100,trans_stress_large=100,
                                 design_shear_stress = 10):
        '''
        Simple buckling calculations according to DNV-RP-C201
        This method is currently not used.
        :return:
        '''

        #7.2 Forces in the idealised stiffened plate

        s = self.spacing
        t = self.plate_th
        l = self.span
        E = 2.1e11

        pSd = design_lat_press*1000
        tauSd = design_shear_stress*1e6
        sigy2Sd =trans_stress_small*1e6
        fy = self.mat_yield

        #7.3 Effective plate width
        alphac = 1.1*(s/t)*math.sqrt(fy/E)
        gamma = 0.21*(alphac-0.2)

        if alphac<=0.2: kappa = 1
        elif 0.2<alphac<2: kappa = (1/(2*(alphac**2)))*(1-+gamma+alphac**2-math.sqrt((1+gamma+alphac**2)**2-4*alphac**2))
        else: kappa=(1/(2*alphac**2))+0.7

        ha = 0.05*(s/t)-0.75
        assert ha>= 0,'ha must be larger than 0'
        kp = 1 if pSd<=2*((t/s)**2)*fy else 1-ha*((pSd/fy)-2*(t/s)**2)

        sigyR=( (1.3*t/l)*math.sqrt(E/fy)+kappa*(1-(1.3*t/l)*math.sqrt(E/fy)))*fy*kp
        sigyRd = sigyR / 1.15

        # plate resistance check
        ksp = math.sqrt(1-3*(tauSd/(fy/1))**2)
        eq7_19 = ksp*sigyRd/sigy2Sd
        return eq7_19

    def buckling_local_stiffener(self):
        '''
        Local requirements for stiffeners. Chapter 9.11.
        :return:
        '''

        epsilon = math.sqrt(235 / (self.mat_yield / 1e6))

        if self.stiffener_type == 'L':
            c = self.flange_width - self.web_th/2
        elif self.stiffener_type == 'T':
            c = self.flange_width/2 - self.web_th/2
        elif self.stiffener_type == 'FB':
            return self.web_height <= 42 * self.web_th * epsilon, self.web_height/(42 * self.web_th * epsilon)

        # print(self.web_height, self.web_th, self.flange_width ,self.flange_th )
        # print('c:',c, 14 * self.flange_th * epsilon, ' | ',  self.web_height, 42 * self.web_th * epsilon)
        # print(c <= (14  * self.flange_th * epsilon) and self.web_height <= 42 * self.web_th * epsilon)
        # print(c/(14  * self.flange_th * epsilon), self.web_height / (42 * self.web_th * epsilon))
        # print('')

        return c <= (14  * self.flange_th * epsilon) and self.web_height <= 42 * self.web_th * epsilon, \
               max(c/(14  * self.flange_th * epsilon), self.web_height / (42 * self.web_th * epsilon))

    def is_acceptable_pl_thk(self, design_pressure):
        '''
        Checking if the thickness is acceptable.
        :return:
        '''
        return self.get_dnv_min_thickness(design_pressure) <= self.plate_th*1000

class CalcFatigue(Structure):
    '''
    This Class does the calculations for the plate fields. 
    Input is a structure object (getters from the Structure Class)
    '''
    def __init__(self, main_dict, fatigue_dict):
        super(CalcFatigue, self).__init__(main_dict, fatigue_dict)
        if fatigue_dict is not None:
            self._sn_curve = fatigue_dict['SN-curve']
            self._acc = fatigue_dict['Accelerations']
            self._weibull = fatigue_dict['Weibull']
            self._period = fatigue_dict['Period']
            self._k_factor = fatigue_dict['SCF']
            self._corr_loc = fatigue_dict['CorrLoc']
            self._no_of_cycles = fatigue_dict['n0']
            self._design_life = fatigue_dict['Design life']
            self._fraction = fatigue_dict['Fraction']
            self._case_order = fatigue_dict['Order']
            try:
                self._dff = fatigue_dict['DFF']
            except KeyError:
                self._dff = 2

            self.fatigue_dict = fatigue_dict

    def get_sn_curve(self):
        return self._sn_curve

    def __get_sigma_ext(self, int_press):
        return -0.5*int_press* ((self.spacing / (self.plate_th))**2) * (self._k_factor/1000**2)

    def __get_sigma_int(self, ext_press):
        return 0.5*ext_press*((self.spacing/(self.plate_th))**2) * (self._k_factor/1000**2)

    def __get_range(self, idx, int_press, ext_press):
        return 2*math.sqrt(math.pow(self.__get_sigma_ext(ext_press), 2) +
                           math.pow(self.__get_sigma_int(int_press), 2) +
                           2*self._corr_loc[idx]*self.__get_sigma_ext(ext_press)
                           *self.__get_sigma_int(int_press))

    def __get_stress_fraction(self,idx, int_press, ext_press):
        return self.__get_range(idx, int_press, ext_press) / \
               math.pow(math.log(self._no_of_cycles), 1/self._weibull[idx])

    def __get_gamma1(self,idx):
        return math.exp(gammaln(snc.get_paramter(self._sn_curve,'m1')/self._weibull[idx] + 1))

    def __get_gamma2(self,idx):
        return math.exp(gammaln(snc.get_paramter(self._sn_curve, 'm2') / self._weibull[idx] + 1))

    def get_damage_slope1(self, idx, curve, int_press=0, ext_press=0):
        m1, log_a1, k, slope = snc.get_paramter(curve,'m1'), snc.get_paramter(curve,'log a1'),\
                               snc.get_paramter(curve,'k'), snc.get_paramter(curve,'slope')
        cycles = self._design_life*365*24*3600/self._period[idx]
        thk_eff = math.log10(max(1,self.plate_th/0.025)) * k
        slope_ch = math.exp( math.log( math.pow(10, log_a1-m1*thk_eff)/slope) / m1)
        gamma1 = self.__get_gamma1(idx)
        weibull = self._weibull[idx]
        stress_frac = self.__get_stress_fraction(idx, int_press, ext_press)
        # print('Internal pressure: ', int_press)
        # print('External pressure: ', ext_press)
        # finding GAMMADIST
        if stress_frac == 0:
            return 0

        x, alpha = math.pow(slope_ch/stress_frac, weibull),1 + m1/weibull
        gamma_val = gammadist.cdf(x,alpha)

        return cycles / math.pow(10, log_a1-m1*thk_eff) * math.pow(stress_frac, m1)*gamma1*(1-gamma_val)\
               *self._fraction[idx]

    def get_damage_slope2(self, idx, curve, int_press, ext_press):
        m2, log_m2, k, slope = snc.get_paramter(curve,'m2'), snc.get_paramter(curve,'log a2'),\
                               snc.get_paramter(curve,'k'), snc.get_paramter(curve,'slope')
        cycles = self._design_life*365*24*3600/self._period[idx]
        thk_eff = math.log10(max(1,self.plate_th/25)) * k
        slope_ch = math.exp( math.log( math.pow(10, log_m2-m2*thk_eff)/slope) / m2)
        gammm2 = self.__get_gamma2(idx)
        weibull = self._weibull[idx]
        stress_frac = self.__get_stress_fraction(idx, int_press, ext_press)

        # finding GAMMADIST
        if stress_frac == 0:
            return 0
        x, alpha = math.pow(slope_ch/stress_frac, weibull),1 + m2/weibull
        gamma_val = gammadist.cdf(x,alpha)

        return cycles / math.pow(10, log_m2-m2*thk_eff) * math.pow(stress_frac, m2)*gammm2*(gamma_val)\
               *self._fraction[idx]

    def get_total_damage(self, int_press=(0, 0, 0), ext_press=(0, 0, 0)):
        damage = 0

        for idx in range(3):
            if self._fraction[idx] != 0 and self._period[idx] != 0:
                damage += self.get_damage_slope1(idx,self._sn_curve, int_press[idx], ext_press[idx]) + \
                          self.get_damage_slope2(idx,self._sn_curve, int_press[idx], ext_press[idx])
        return damage

    def set_fatigue_properties(self, fatigue_dict):
        ''' Setting the fatiuge properties. '''
        self._sn_curve, self.fatigue_dict['SN-curve'] = fatigue_dict['SN-curve'], fatigue_dict['SN-curve']
        self._acc, self.fatigue_dict['Accelerations'] = fatigue_dict['Accelerations'], fatigue_dict['Accelerations']
        self._weibull, self.fatigue_dict['Weibull'] = fatigue_dict['Weibull'], fatigue_dict['Weibull']
        self._period, self.fatigue_dict['Period'] = fatigue_dict['Period'], fatigue_dict['Period']
        self._k_factor, self.fatigue_dict['SCF'] = fatigue_dict['SCF'], fatigue_dict['SCF']
        self._corr_loc, self.fatigue_dict['CorrLoc'] = fatigue_dict['CorrLoc'], fatigue_dict['CorrLoc']
        self._no_of_cycles, self.fatigue_dict['n0'] = fatigue_dict['n0'], fatigue_dict['n0']
        self._design_life, self.fatigue_dict['Design life'] = fatigue_dict['Design life'], fatigue_dict['Design life']
        self._fraction, self.fatigue_dict['Fraction'] = fatigue_dict['Fraction'], fatigue_dict['Fraction']
        self._case_order, self.fatigue_dict['Order'] = fatigue_dict['Order'], fatigue_dict['Order']
        self._dff, self.fatigue_dict['DFF'] = fatigue_dict['DFF'], fatigue_dict['DFF']

    def get_fatigue_properties(self):
        ''' Returning properties as a dictionary '''
        return self.fatigue_dict

    def get_accelerations(self):
        ''' Returning tuple of accelerattions.'''
        return self._acc

    def get_dff(self):
        return self._dff

    def get_design_life(self):
        return self._design_life

if __name__ == '__main__':
    import ANYstructure.example_data as test

    # print('Fatigue test: ')
    # my_test = CalcFatigue(test.obj_dict, test.fat_obj_dict)
    # print('Total damage: ',my_test.get_total_damage(int_press=(0,0,0), ext_press=(50000, 60000,0)))
    # print('')
    # print('Buckling test: ')
    #
    # my_buc = test.get_structure_calc_object()
    #
    # #print(my_buc.calculate_buckling_all(design_lat_press=100))
    # print(my_buc.calculate_slamming_plate(1000000))
    # print(my_buc.calculate_slamming_stiffener(1000000))
    # print(my_buc.get_net_effective_plastic_section_modulus())

    #my_test.get_total_damage(int_press=(0, 0, 0), ext_press=(0, 40000, 0))
    import ANYstructure.example_data as ex
    for example in [CalcScantlings(ex.obj_dict), CalcScantlings(ex.obj_dict2), CalcScantlings(ex.obj_dict_L)]:
        my_test = example
        # my_test = CalcScantlings(example)
        # my_test = CalcFatigue(example, test.fat_obj_dict2)
        # my_test.get_total_damage(int_press=(0, 0, 0), ext_press=(0, 40000, 0))
        # print('Total damage: ', my_test.get_total_damage(int_press=(0, 0, 0), ext_press=(0, 40000, 0)))
        # print(my_test.get_fatigue_properties())
        pressure = 200
        print(my_test.buckling_local_stiffener())
        # print('SHEAR CENTER: ',my_test.get_shear_center())
        # print('SECTION MOD: ',my_test.get_section_modulus())
        # print('SECTION MOD FLANGE: ', my_test.get_section_modulus()[0])
        # print('SHEAR AREA: ', my_test.get_shear_area())
        # print('PLASTIC SECTION MOD: ',my_test.get_plasic_section_modulus())
        # print('MOMENT OF INTERTIA: ',my_test.get_moment_of_intertia())
        # print('WEIGHT', my_test.get_weight())
        # print('PROPERTIES', my_test.get_structure_prop())
        # print('CROSS AREA', my_test.get_cross_section_area())
        # print()
        #
        # print('EFFICIENT MOMENT OF INTERTIA: ',my_test.get_moment_of_intertia(efficent_se=my_test.get_plate_efficent_b(
        #     design_lat_press=pressure)))
        # print('Se: ',my_test.calculate_buckling_all(design_lat_press=pressure,checked_side='s'))
        # print('Se: ', my_test.calculate_buckling_all(design_lat_press=pressure, checked_side='p'))
        # print('MINIMUM PLATE THICKNESS',my_test.get_dnv_min_thickness(pressure))
        # print('MINIMUM SECTION MOD.', my_test.get_dnv_min_section_modulus(pressure))
        # print()
