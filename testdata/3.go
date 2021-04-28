package main

func __init__() {
	super(Structure, self).__init__()
	self.main_dict = main_dict
	self.plate_th = main_dict["plate_thk"][0]
	self.web_height = main_dict["stf_web_height"][0]
	self.web_th = main_dict["stf_web_thk"][0]
	self.flange_width = main_dict["stf_flange_width"][0]
	self.flange_th = main_dict["stf_flange_thk"][0]
	self.mat_yield = main_dict["mat_yield"][0]
	self.span = main_dict["span"][0]
	self.spacing = main_dict["spacing"][0]
	self.structure_type = main_dict["structure_type"][0]
	self.sigma_y1 = main_dict["sigma_y1"][0]
	self.sigma_y2 = main_dict["sigma_y2"][0]
	self.sigma_x = main_dict["sigma_x"][0]
	self.tauxy = main_dict["tau_xy"][0]
	self.plate_kpp = main_dict["plate_kpp"][0]
	self.stf_kps = main_dict["stf_kps"][0]
	self.km1 = main_dict["stf_km1"][0]
	self.km2 = main_dict["stf_km2"][0]
	self.km3 = main_dict["stf_km3"][0]
	self.stiffener_type = main_dict["stf_type"][0]
	self.sigma_y = self.sigma_y2 + (self.sigma_y1-self.sigma_y2)*(min(0.25*self.span, 0.5*self.spacing)/self.span)
}
func __str__() {
	"\n        Returning all properties.\n        "
	return str("\n Plate field span:              " + str(round(self.span, 1)) + " meters" + "\n Stiffener spacing:             " + str(self.spacing*1000) + " mm" + "\n Plate thickness:               " + str(self.plate_th*1000) + " mm" + "\n Stiffener web height:          " + str(self.web_height*1000) + " mm" + "\n Stiffener web thickness:       " + str(self.web_th*1000) + " mm" + "\n Stiffener flange width:        " + str(self.flange_width*1000) + " mm" + "\n Stiffener flange thickness:    " + str(self.flange_th*1000) + " mm" + "\n Material yield:                " + str(self.mat_yield/1000000.0) + " MPa" + "\n Structure type/stiffener type: " + str(self.structure_type) + "/" + self.stiffener_type + "\n Plate fixation paramter,kpp:   " + str(self.plate_kpp) + " " + "\n Stf. fixation paramter,kps:    " + str(self.stf_kps) + " " + "\n Global stress, sig_y1/sig_y2:  " + str(round(self.sigma_y1, 1)) + "/" + str(round(self.sigma_y2, 1)) + " MPa" + "\n Global stress, sig_x:          " + str(round(self.sigma_x, 1)) + " MPa" + "\n Global shear, tau_xy:          " + str(round(self.tauxy, 1)) + " MPa" + "\n km1,km2,km3:                   " + str(self.km1) + "/" + str(self.km2) + "/" + str(self.km3) + "\n Pressure side (p-plate/s-stf): " + str(self.pressure_side) + " ")
}
func get_one_line_string() {
	" Returning a one line string. "
	return "pl_" + str(round(self.spacing*1000, 1)) + "x" + str(round(self.plate_th*1000, 1)) + " stf_" + self.stiffener_type + str(round(self.web_height*1000, 1)) + "x" + str(round(self.web_th*1000, 1)) + "+" + str(round(self.flange_width*1000, 1)) + "x" + str(round(self.flange_th*1000, 1))
}
func get_report_stresses() {
	"Return the stresses to the report"
	return "sigma_y1: " + str(round(self.sigma_y1, 1)) + " sigma_y2: " + str(round(self.sigma_y2, 1)) + " sigma_x: " + str(round(self.sigma_x, 1)) + " tauxy: " + str(round(self.tauxy, 1))
}
func get_extended_string() {
	" Some more information returned. "
	return "span: " + str(round(self.span, 4)) + " structure type: " + self.structure_type + " stf. type: " + self.stiffener_type + " pressure side: " + self.pressure_side
}
func get_sigma_y1() {
	"\n        Return sigma_y1\n        :return:\n        "
	return self.sigma_y1
}
func get_sigma_y2() {
	"\n        Return sigma_y2\n        :return:\n        "
	return self.sigma_y1
}
func get_sigma_x() {
	"\n        Return sigma_x\n        :return:\n        "
	return self.sigma_y1
}
func get_tau_xy() {
	"\n        Return tau_xy\n        :return:\n        "
	return self.tauxy
}
func get_s() {
	"\n        Return the spacing\n        :return:\n        "
	return self.spacing
}
func get_pl_thk() {
	"\n        Return the plate thickness\n        :return:\n        "
	return self.plate_th
}
func get_web_h() {
	"\n        Return the web heigh\n        :return:\n        "
	return self.web_height
}
func get_web_thk() {
	"\n        Return the spacing\n        :return:\n        "
	return self.web_th
}
func get_fl_w() {
	"\n        Return the flange width\n        :return:\n        "
	return self.flange_width
}
func get_fl_thk() {
	"\n        Return the flange thickness\n        :return:\n        "
	return self.flange_th
}
func get_fy() {
	"\n        Return material yield\n        :return:\n        "
	return self.mat_yield
}
func get_span() {
	"\n        Return the span\n        :return:\n        "
	return self.span
}
func get_lg() {
	"\n        Return the girder length\n        :return:\n        "
	return self.girder_lg
}
func get_kpp() {
	"\n        Return var\n        :return:\n        "
	return self.plate_kpp
}
func get_kps() {
	"\n        Return var\n        :return:\n        "
	return self.stf_kps
}
func get_km1() {
	"\n        Return var\n        :return:\n        "
	return self.km1
}
func get_km2() {
	"\n        Return var\n        :return:\n        "
	return self.km2
}
func get_km3() {
	"\n        Return var\n        :return:\n        "
	return self.km3
}
func get_side() {
	"\n        Return the checked pressure side.\n        :return: \n        "
	return self.pressure_side
}
func get_tuple() {
	" Return a tuple of the plate stiffener"
	return self.spacing.self.plate_th.self.web_height.self.web_th.self.flange_width.self.flange_th.self.span.self.girder_lg.self.stiffener_type
}
func get_section_modulus() {
	"\n        Returns the section modulus.\n        :param efficient_se: \n        :return: \n        "
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'efficient_se'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       Eq (
	//
	//       ) // Eq
	//       ] //
	//       comparators =  [
	//       Name (
	//       id = 'None'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ] //
	//       ) // Compare
	//       body = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'spacing'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       orelse = Name (
	//       id = 'efficient_se'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // IfExp
	//
	tf1 = self.plate_th
	tf2 = self.flange_th
	b2 = self.flange_width
	h = self.flange_th + self.web_height + self.plate_th
	tw = self.web_th
	hw = self.web_height
	Ax = tf1*b1 + tf2*b2 + hw*tw
	ez = (tf1*b1*tf1/2 + hw*tw*(tf1+hw/2) + tf2*b2*(tf1+hw+tf2/2)) / Ax
	Iyc = 1 / 12 * (b1*math.pow(tf1, 3) + b2*math.pow(tf2, 3) + tw*math.pow(hw, 3))
	Iy = Iyc + (tf1*b1*math.pow(tf1/2, 2) + tw*hw*math.pow(tf1+hw/2, 2) + tf2*b2*math.pow(tf1+hw+tf2/2, 2)) - Ax*math.pow(ez, 2)
	Wey1 = Iy / (h - ez)
	Wey2 = Iy / ez
	return Wey1.Wey2
}
func get_plasic_section_modulus() {
	"\n        Returns the plastic section modulus\n        :return:\n        "
	tf1 = self.plate_th
	tf2 = self.flange_th
	b1 = self.spacing
	b2 = self.flange_width
	h = self.flange_th + self.web_height + self.plate_th
	tw = self.web_th
	hw = self.web_height
	Ax = tf1*b1 + tf2*b2 + (h-tf1-tf2)*tw
	ezpl = (Ax/2-b1*tf1)/tw + tf1
	az1 = h - ezpl - tf1
	az2 = ezpl - tf2
	Wy1 = b1*tf1*(az1+tf1/2) + tw/2*math.pow(az1, 2)
	Wy2 = b2*tf2*(az2+tf2/2) + tw/2*math.pow(az2, 2)
	return Wy1 + Wy2
}
func get_shear_center() {
	"\n        Returning the shear center\n        :return:\n        "
	tf1 = self.plate_th
	tf2 = self.flange_th
	b1 = self.spacing
	b2 = self.flange_width
	h = self.flange_th + self.web_height + self.plate_th
	tw = self.web_th
	hw = self.web_height
	Ax = tf1*b1 + tf2*b2 + (h-tf1-tf2)*tw
	ez = (b2*tf2*tf2/2 + tw*hw*(tf2+hw/2) + tf1*b1*(tf2+hw+tf1/2)) / Ax
	Iz1 = tf1 * math.pow(b1, 3)
	Iz2 = tf2 * math.pow(b2, 3)
	ht = h - tf1/2 - tf2/2
	return Iz1*ht/(Iz1+Iz2) + tf2/2 - ez
}
func get_moment_of_intertia() {
	"\n        Returning moment of intertia.\n        :return:\n        "
	tf1 = self.plate_th
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'efficent_se'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       Eq (
	//
	//       ) // Eq
	//       ] //
	//       comparators =  [
	//       Name (
	//       id = 'None'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ] //
	//       ) // Compare
	//       body = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'spacing'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       orelse = Name (
	//       id = 'efficent_se'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // IfExp
	//
	h = self.flange_th + self.web_height + self.plate_th
	tw = self.web_th
	hw = self.web_height
	tf2 = self.flange_th
	b2 = self.flange_width
	Ax = tf1*b1 + tf2*b2 + (h-tf1-tf2)*tw
	Iyc = 1 / 12 * (b1*math.pow(tf1, 3) + b2*math.pow(tf2, 3) + tw*math.pow(hw, 3))
	ez = (tf1*b1*(h-tf1/2) + hw*tw*(tf2+hw/2) + tf2*b2*(tf2/2)) / Ax
	Iy = Iyc + (tf1*b1*math.pow(tf2+hw+tf1/2, 2) + tw*hw*math.pow(tf2+hw/2, 2) + tf2*b2*math.pow(tf2/2, 2)) - Ax*math.pow(ez, 2)
	return Iy
}
func get_structure_prop() {
	return self.main_dict
}
func get_structure_type() {
	return self.structure_type
}
func get_stiffener_type() {
	return self.stiffener_type
}
func get_shear_area() {
	"\n        Returning the shear area in [m^2]\n        :return:\n        "
	return self.flange_th*self.web_th + self.web_th*self.plate_th + self.web_height*self.web_th
}
func set_main_properties() {
	"\n        Resettting all properties\n        :param input_dictionary:\n        :return:\n        "
	self.main_dict = main_dict
	self.plate_th = main_dict["plate_thk"][0]
	self.web_height = main_dict["stf_web_height"][0]
	self.web_th = main_dict["stf_web_thk"][0]
	self.flange_width = main_dict["stf_flange_width"][0]
	self.flange_th = main_dict["stf_flange_thk"][0]
	self.mat_yield = main_dict["mat_yield"][0]
	self.span = main_dict["span"][0]
	self.spacing = main_dict["spacing"][0]
	self.structure_type = main_dict["structure_type"][0]
	self.sigma_y1 = main_dict["sigma_y1"][0]
	self.sigma_y2 = main_dict["sigma_y2"][0]
	self.sigma_x = main_dict["sigma_x"][0]
	self.tauxy = main_dict["tau_xy"][0]
	self.plate_kpp = main_dict["plate_kpp"][0]
	self.stf_kps = main_dict["stf_kps"][0]
	self.km1 = main_dict["stf_km1"][0]
	self.km2 = main_dict["stf_km2"][0]
	self.km3 = main_dict["stf_km3"][0]
	self.stiffener_type = main_dict["stf_type"][0]
}
func set_stresses() {
	"\n        Setting the global stresses.\n        :param sigy1:\n        :param sigy2:\n        :param sigx:\n        :param tauxy:\n        :return:\n        "
	self.main_dict["sigma_y1"][0] = sigy1
	self.sigma_y1 = sigy1
	self.main_dict["sigma_y2"][0] = sigy2
	self.sigma_y2 = sigy2
	self.main_dict["sigma_x"][0] = sigx
	self.sigma_x = sigx
	self.main_dict["tau_xy"][0] = tauxy
	self.tauxy = tauxy
}
func get_plate_thk() {
	"\n        Return the plate thickness\n        :return:\n        "
	return self.plate_th
}
func get_cross_section_area() {
	"\n        Returns the cross section area.\n        :return:\n        "
	tf1 = self.plate_th
	tf2 = self.flange_th
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'efficient_se'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       Eq (
	//
	//       ) // Eq
	//       ] //
	//       comparators =  [
	//       Name (
	//       id = 'None'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ] //
	//       ) // Compare
	//       body = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'spacing'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       orelse = Name (
	//       id = 'efficient_se'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // IfExp
	//
	b2 = self.flange_width
	h = self.flange_th + self.web_height + self.plate_th
	tw = self.web_th
	return tf1*b1 + tf2*b2 + (h-tf1-tf2)*tw
}
func get_cross_section_centroid_with_effective_plate() {
	"\n        Returns cross section centroid\n        :return:\n        "
	tf1 = self.plate_th
	tf2 = self.flange_th
	b1 = se
	b2 = self.flange_width
	h = self.flange_th + self.web_height + self.plate_th
	tw = self.web_th
	hw = self.web_height
	Ax = tf1*b1 + tf2*b2 + hw*tw
	return (tf1*b1*tf1/2 + hw*tw*(tf1+hw/2) + tf2*b2*(tf1+hw+tf2/2)) / Ax
}
func get_weight() {
	"\n        Return the weight.\n        :return:\n        "
	return 7850 * self.span * (self.spacing*self.plate_th + self.web_height*self.web_th + self.flange_width*self.flange_th)
}
func get_weight_width_lg() {
	"\n        Return the weight including Lg\n        :return:\n        "
	pl_area = self.girder_lg * self.plate_th
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in func transpileExprs
	//    │  ├──Error in Operation: BinOp (
	//    │  │  left = Attribute (
	//    │  │  value = Name (
	//    │  │  id = 'self'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  attr = 'girder_lg'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Attribute
	//    │  │  op = FloorDiv (
	//    │  │
	//    │  │  ) // FloorDiv
	//    │  │  right = Attribute (
	//    │  │  value = Name (
	//    │  │  id = 'self'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  attr = 'spacing'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Attribute
	//    │  │  ) // BinOp
	//    │  │  Error in func transpileOp
	//    │  │  └──not valid token: FloorDiv for FloorDiv (
	//    │  │
	//    │  │  ) // FloorDiv
	//    │  └──Not valid BinaryOp: BinOp (
	//    │     left = Attribute (
	//    │     value = Name (
	//    │     id = 'self'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     attr = 'girder_lg'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Attribute
	//    │     op = FloorDiv (
	//    │
	//    │     ) // FloorDiv
	//    │     right = Attribute (
	//    │     value = Name (
	//    │     id = 'self'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     attr = 'spacing'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Attribute
	//    │     ) // BinOp
	//    └──Not valid BinaryOp: BinOp (
	//       left = BinOp (
	//       left = BinOp (
	//       left = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'web_height'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'web_th'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       ) // BinOp
	//       op = Add (
	//
	//       ) // Add
	//       right = BinOp (
	//       left = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'flange_width'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'flange_th'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       ) // BinOp
	//       ) // BinOp
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'girder_lg'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       op = FloorDiv (
	//
	//       ) // FloorDiv
	//       right = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'spacing'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       ) // BinOp
	//       ) // BinOp
	//
	return (pl_area + stf_area) * 7850 * self.span
}
func set_span() {
	"\n        Setting the span. Used when moving a point.\n        :return: \n        "
	self.span = span
	self.main_dict["span"][0] = span
}
func __init__() {
	super(CalcScantlings, self).__init__()
	self.lat_press = lat_press
	self.category = category
	self._need_recalc = True
}
func need_recalc() {
	return self._need_recalc
}
func need_recalc() {
	self._need_recalc = val
}
func get_results_for_report() {
	"\n        Returns a string for the report.\n        :return:\n        "
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: ListComp
	//       cannot transpile : ListComp (
	//       elt = Call (
	//       func = Name (
	//       id = 'round'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       args =  [
	//       Name (
	//       id = 'res'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       Num (
	//       n = 1
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       generators =  [
	//       comprehension (
	//       target = Name (
	//       id = 'res'
	//       ctx = Store (
	//
	//       ) // Store
	//       ) // Name
	//       iter = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'calculate_buckling_all'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//
	//       ] //
	//       keywords =  [
	//       keyword (
	//       arg = 'design_lat_press'
	//       value = Name (
	//       id = 'lat_press'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // keyword
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ifs =  [
	//
	//       ] //
	//       ) // comprehension
	//       ] //
	//       ) // ListComp
	//
	return "Minimum section modulus:" + str(int(self.get_dnv_min_section_modulus()*math.Pow(1000, 3))) + "mm^3 " + " Minium plate thickness: " + str(round(self.get_dnv_min_thickness(), 1)) + " Buckling results: eq7_19: " + str(buc[0]) + " eq7_50: " + str(buc[1]) + " eq7_51: " + str(buc[2]) + " eq7_52: " + str(buc[3]) + " eq7_53: " + str(buc[4])
}
func calculate_slamming_plate() {
	" Slamming pressure input is Pa "
	ka1 = 1.1
	ka2 = min(max(0.4, self.spacing/self.span), 1)
	ka = math.pow(ka1-0.25*ka2, 2)
	sigmaf = self.mat_yield / 1000000.0
	psl = slamming_pressure / 1000
	Cd = 1.5
	return 0.0158 * ka * self.spacing * 1000 * math.sqrt(psl/(Cd*sigmaf))
}
func calculate_slamming_stiffener() {
	tk = 0
	psl = slamming_pressure / 1000
	Pst = psl / 2
	sigmaf = self.mat_yield / 1000000.0
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: ListComp
	//       cannot transpile : ListComp (
	//       elt = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'val'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Sub (
	//
	//       ) // Sub
	//       right = Name (
	//       id = 'tk'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Num (
	//       n = 1000
	//       ) // Num
	//       ) // BinOp
	//       generators =  [
	//       comprehension (
	//       target = Name (
	//       id = 'val'
	//       ctx = Store (
	//
	//       ) // Store
	//       ) // Name
	//       iter = List (
	//       elts =  [
	//       Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'web_height'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'web_th'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'plate_th'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'flange_th'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'flange_width'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'spacing'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       ] //
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // List
	//       ifs =  [
	//
	//       ] //
	//       ) // comprehension
	//       ] //
	//       ) // ListComp
	//
	ns = 2
	tau_eH = sigmaf / math.sqrt(3)
	h_stf = (self.web_height + self.flange_th) * 1000
	f_shr = 0.7
	lbdg = self.span
	lshr = self.span - self.spacing/4000
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Num (
	//       n = 75
	//       ) // Num
	//       ops =  [
	//       LtE (
	//
	//       ) // LtE
	//       LtE (
	//
	//       ) // LtE
	//       ] //
	//       comparators =  [
	//       Name (
	//       id = 'angle'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       Num (
	//       n = 90
	//       ) // Num
	//       ] //
	//       ) // Compare
	//       body = BinOp (
	//       left = Name (
	//       id = 'h_stf'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Add (
	//
	//       ) // Add
	//       right = Name (
	//       id = 'tp'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       orelse = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'h_stf'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Add (
	//
	//       ) // Add
	//       right = Name (
	//       id = 'tp'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'sin'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'radians'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       Name (
	//       id = 'angle'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       ) // IfExp
	//
	tw = f_shr * Pst * s * lshr / (dshr * tau_eH)
	if self.web_th*1000 < tw {
		// Find PY4GO error
		// Error in func transpileStmt
		// └──Error in func transpileExprs
		//    ├──Error in func transpileExprs
		//    │  ├──Error function: 5
		//    │  └──Name: Dict
		//    │     cannot transpile : Dict (
		//    │     keys =  [
		//    │     Str (
		//    │     s = 'tw_req'
		//    │     ) // Str
		//    │     Str (
		//    │     s = 'Zp_req'
		//    │     ) // Str
		//    │     ] //
		//    │     values =  [
		//    │     Name (
		//    │     id = 'tw'
		//    │     ctx = Load (
		//    │
		//    │     ) // Load
		//    │     ) // Name
		//    │     Name (
		//    │     id = 'None'
		//    │     ctx = Load (
		//    │
		//    │     ) // Load
		//    │     ) // Name
		//    │     ] //
		//    │     ) // Dict
		//    ├──Error function: 0
		//    └──Error in func transpileExprs
		//       ├──Error function: 5
		//       └──Name: Dict
		//          cannot transpile : Dict (
		//          keys =  [
		//          Str (
		//          s = 'tw_req'
		//          ) // Str
		//          Str (
		//          s = 'Zp_req'
		//          ) // Str
		//          ] //
		//          values =  [
		//          Name (
		//          id = 'tw'
		//          ctx = Load (
		//
		//          ) // Load
		//          ) // Name
		//          Name (
		//          id = 'None'
		//          ctx = Load (
		//
		//          ) // Load
		//          ) // Name
		//          ] //
		//          ) // Dict
		//
	}
	fpl = 8 * (1 + ns/2)
	Zp_req = 1.2*Pst*s*math.pow(lbdg, 2)/(fpl*sigmaf) + ns*(1-math.sqrt(1-math.pow(tw/twa, 2)))*hw*tw*(hw+tp)/8000
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in func transpileExprs
	//    │  ├──Error function: 5
	//    │  └──Name: Dict
	//    │     cannot transpile : Dict (
	//    │     keys =  [
	//    │     Str (
	//    │     s = 'tw_req'
	//    │     ) // Str
	//    │     Str (
	//    │     s = 'Zp_req'
	//    │     ) // Str
	//    │     ] //
	//    │     values =  [
	//    │     Name (
	//    │     id = 'tw'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     Name (
	//    │     id = 'Zp_req'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     ] //
	//    │     ) // Dict
	//    ├──Error function: 0
	//    └──Error in func transpileExprs
	//       ├──Error function: 5
	//       └──Name: Dict
	//          cannot transpile : Dict (
	//          keys =  [
	//          Str (
	//          s = 'tw_req'
	//          ) // Str
	//          Str (
	//          s = 'Zp_req'
	//          ) // Str
	//          ] //
	//          values =  [
	//          Name (
	//          id = 'tw'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          Name (
	//          id = 'Zp_req'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          ] //
	//          ) // Dict
	//
}
func check_all_slamming() {
	" A summary check of slamming "
	pl_chk = self.calculate_slamming_plate(slamming_pressure)
	if self.plate_th*1000 < pl_chk {
		chk1 = pl_chk / self.plate_th * 1000
		return False.chk1
	}
	stf_res = self.calculate_slamming_stiffener(slamming_pressure)
	if self.web_th*1000 < stf_res["tw_req"] {
		chk2 = stf_res["tw_req"] / self.web_th * 1000
		return False.chk2
	}
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in Operation: Compare (
	//    │  left = Subscript (
	//    │  value = Name (
	//    │  id = 'stf_res'
	//    │  ctx = Load (
	//    │
	//    │  ) // Load
	//    │  ) // Name
	//    │  slice = Index (
	//    │  value = Str (
	//    │  s = 'Zp_req'
	//    │  ) // Str
	//    │  ) // Index
	//    │  ctx = Load (
	//    │
	//    │  ) // Load
	//    │  ) // Subscript
	//    │  ops =  [
	//    │  IsNot (
	//    │
	//    │  ) // IsNot
	//    │  ] //
	//    │  comparators =  [
	//    │  Name (
	//    │  id = 'None'
	//    │  ctx = Load (
	//    │
	//    │  ) // Load
	//    │  ) // Name
	//    │  ] //
	//    │  ) // Compare
	//    │  Error in func transpileOp
	//    │  └──not valid token: IsNot for IsNot (
	//    │
	//    │  ) // IsNot
	//    └──Nil binary: Compare (
	//       left = Subscript (
	//       value = Name (
	//       id = 'stf_res'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       slice = Index (
	//       value = Str (
	//       s = 'Zp_req'
	//       ) // Str
	//       ) // Index
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Subscript
	//       ops =  [
	//       IsNot (
	//
	//       ) // IsNot
	//       ] //
	//       comparators =  [
	//       Name (
	//       id = 'None'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ] //
	//       ) // Compare
	//
	return True.None
}
func get_net_effective_plastic_section_modulus() {
	" Calculated according to Rules for classification: Ships \xe2\x80\x94 DNVGL-RU-SHIP Pt.3 Ch.3. Edition July 2017,\n            page 83 "
	tk = 0
	angle_rad = math.radians(angle)
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: ListComp
	//       cannot transpile : ListComp (
	//       elt = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'val'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Sub (
	//
	//       ) // Sub
	//       right = Name (
	//       id = 'tk'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Num (
	//       n = 1000
	//       ) // Num
	//       ) // BinOp
	//       generators =  [
	//       comprehension (
	//       target = Name (
	//       id = 'val'
	//       ctx = Store (
	//
	//       ) // Store
	//       ) // Name
	//       iter = List (
	//       elts =  [
	//       Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'web_height'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'web_th'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'plate_th'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'flange_th'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'flange_width'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       ] //
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // List
	//       ifs =  [
	//
	//       ] //
	//       ) // comprehension
	//       ] //
	//       ) // ListComp
	//
	h_stf = (self.web_height + self.flange_th) * 1000
	de_gr = 0
	tw_gr = self.web_th * 1000
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'get_stiffener_type'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ops =  [
	//       NotEq (
	//
	//       ) // NotEq
	//       ] //
	//       comparators =  [
	//       Str (
	//       s = 'L'
	//       ) // Str
	//       ] //
	//       ) // Compare
	//       body = BinOp (
	//       left = Name (
	//       id = 'h_stf'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Num (
	//       n = 0.5
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'tf'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ) // BinOp
	//       orelse = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'h_stf'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Sub (
	//
	//       ) // Sub
	//       right = Name (
	//       id = 'de_gr'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Num (
	//       n = 0.5
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'tf'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ) // BinOp
	//       ) // IfExp
	//
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'get_stiffener_type'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ops =  [
	//       Eq (
	//
	//       ) // Eq
	//       ] //
	//       comparators =  [
	//       Str (
	//       s = 'T'
	//       ) // Str
	//       ] //
	//       ) // Compare
	//       body = Num (
	//       n = 0
	//       ) // Num
	//       orelse = BinOp (
	//       left = Num (
	//       n = 0.5
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = Name (
	//       id = 'tf'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Sub (
	//
	//       ) // Sub
	//       right = Name (
	//       id = 'tw_gr'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ) // BinOp
	//       ) // IfExp
	//
	beta = 0.5
	gamma = (1 + math.sqrt(3+12*beta)) / 4
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'get_stiffener_type'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ops =  [
	//       Eq (
	//
	//       ) // Eq
	//       ] //
	//       comparators =  [
	//       Str (
	//       s = 'FB'
	//       ) // Str
	//       ] //
	//       ) // Compare
	//       body = Num (
	//       n = 0
	//       ) // Num
	//       orelse = BinOp (
	//       left = Name (
	//       id = 'bf'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'tf'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ) // IfExp
	//
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in Operation: Compare (
	//    │  left = Num (
	//    │  n = 75
	//    │  ) // Num
	//    │  ops =  [
	//    │  LtE (
	//    │
	//    │  ) // LtE
	//    │  LtE (
	//    │
	//    │  ) // LtE
	//    │  ] //
	//    │  comparators =  [
	//    │  Name (
	//    │  id = 'angle'
	//    │  ctx = Load (
	//    │
	//    │  ) // Load
	//    │  ) // Name
	//    │  Num (
	//    │  n = 90
	//    │  ) // Num
	//    │  ] //
	//    │  ) // Compare
	//    │  Error in func transpileOp
	//    │  └──not valid token: LtE for LtE (
	//    │
	//    │  ) // LtE
	//    └──Nil binary: Compare (
	//       left = Num (
	//       n = 75
	//       ) // Num
	//       ops =  [
	//       LtE (
	//
	//       ) // LtE
	//       LtE (
	//
	//       ) // LtE
	//       ] //
	//       comparators =  [
	//       Name (
	//       id = 'angle'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       Num (
	//       n = 90
	//       ) // Num
	//       ] //
	//       ) // Compare
	//
	return zpl
}
func get_dnv_min_section_modulus() {
	" Section modulus according to DNV rules "
	design_pressure = design_pressure_kpa
	fy = self.mat_yield / 1000000.0
	fyd = fy / 1.15
	sigma_jd = math.sqrt(math.pow(self.sigma_x, 2) + math.pow(self.sigma_y, 2) - self.sigma_x*self.sigma_y + 3*math.pow(self.tauxy, 2))
	sigma_pd2 = fyd - sigma_jd
	kps = self.stf_kps
	km_sides = min(self.km1, self.km3)
	km_middle = self.km2
	Zs = math.pow(self.span, 2) * self.spacing * design_pressure / (min(km_middle, km_sides) * sigma_pd2 * kps) * math.pow(10, 6)
	return max(math.pow(15, 3)/math.pow(1000, 3), Zs/math.pow(1000, 3))
}
func get_dnv_min_thickness() {
	"\n        Return minimum thickness in mm\n        :param design_pressure_kpa:\n        :return:\n        "
	design_pressure = design_pressure_kpa
	sigma_jd = math.sqrt(math.pow(self.sigma_x, 2) + math.pow(self.sigma_y, 2) - self.sigma_x*self.sigma_y + 3*math.pow(self.tauxy, 2))
	fy = self.mat_yield / 1000000
	fyd = fy / 1.15
	sigma_pd1 = min(1.3*(fyd-sigma_jd), fyd)
	sigma_pd1 = abs(sigma_pd1)
	if self.category == "secondary" {
		t0 = 5
	} else {
		t0 = 7
	}
	t_min = 14.3 * t0 / math.sqrt(fyd)
	ka = math.pow(1.1-0.25*self.spacing/self.span, 2)
	if ka > 1 {
		ka = 1
	} else {
		if ka < 0.72 {
			ka = 0.72
		}
	}
	t_min_bend = 15.8 * ka * self.spacing * math.sqrt(design_pressure) / math.sqrt(sigma_pd1*self.plate_kpp)
	if self.lat_press {
		return max(t_min, t_min_bend)
	} else {
		return t_min
	}
}
func get_minimum_shear_area() {
	"\n        Calculating minimum section area according to ch 6.4.4.\n\n        Return [m^2]\n        :return:\n        "
	l = self.span
	s = self.spacing
	fy = self.mat_yield
	fyd = fy / 1.15 / 1000000.0
	sigxd = self.sigma_x
	taupds = 0.577 * math.sqrt(math.pow(fyd, 2)-math.pow(sigxd, 2))
	As = l * s * pressure / (2 * taupds) * math.pow(10, 3)
	return As / math.pow(1000, 2)
}
func is_acceptable_sec_mod() {
	"\n        Checking if the result is accepable.\n        :param section_module:\n        :param pressure:\n        :return:\n        "
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in func transpileExprs
	//    │  ├──Error in Operation: Compare (
	//    │  │  left = Call (
	//    │  │  func = Name (
	//    │  │  id = 'min'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  args =  [
	//    │  │  Name (
	//    │  │  id = 'section_module'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  ] //
	//    │  │  keywords =  [
	//    │  │
	//    │  │  ] //
	//    │  │  starargs = None
	//    │  │  kwargs = None
	//    │  │  ) // Call
	//    │  │  ops =  [
	//    │  │  GtE (
	//    │  │
	//    │  │  ) // GtE
	//    │  │  ] //
	//    │  │  comparators =  [
	//    │  │  Call (
	//    │  │  func = Attribute (
	//    │  │  value = Name (
	//    │  │  id = 'self'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  attr = 'get_dnv_min_section_modulus'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Attribute
	//    │  │  args =  [
	//    │  │  Name (
	//    │  │  id = 'pressure'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  ] //
	//    │  │  keywords =  [
	//    │  │
	//    │  │  ] //
	//    │  │  starargs = None
	//    │  │  kwargs = None
	//    │  │  ) // Call
	//    │  │  ] //
	//    │  │  ) // Compare
	//    │  │  Error in func transpileOp
	//    │  │  └──not valid token: GtE for GtE (
	//    │  │
	//    │  │  ) // GtE
	//    │  └──Nil binary: Compare (
	//    │     left = Call (
	//    │     func = Name (
	//    │     id = 'min'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     args =  [
	//    │     Name (
	//    │     id = 'section_module'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     ] //
	//    │     keywords =  [
	//    │
	//    │     ] //
	//    │     starargs = None
	//    │     kwargs = None
	//    │     ) // Call
	//    │     ops =  [
	//    │     GtE (
	//    │
	//    │     ) // GtE
	//    │     ] //
	//    │     comparators =  [
	//    │     Call (
	//    │     func = Attribute (
	//    │     value = Name (
	//    │     id = 'self'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     attr = 'get_dnv_min_section_modulus'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Attribute
	//    │     args =  [
	//    │     Name (
	//    │     id = 'pressure'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     ] //
	//    │     keywords =  [
	//    │
	//    │     ] //
	//    │     starargs = None
	//    │     kwargs = None
	//    │     ) // Call
	//    │     ] //
	//    │     ) // Compare
	//    ├──Error function: 0
	//    └──Error in func transpileExprs
	//       ├──Error in Operation: Compare (
	//       │  left = Call (
	//       │  func = Name (
	//       │  id = 'min'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  args =  [
	//       │  Name (
	//       │  id = 'section_module'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  ] //
	//       │  keywords =  [
	//       │
	//       │  ] //
	//       │  starargs = None
	//       │  kwargs = None
	//       │  ) // Call
	//       │  ops =  [
	//       │  GtE (
	//       │
	//       │  ) // GtE
	//       │  ] //
	//       │  comparators =  [
	//       │  Call (
	//       │  func = Attribute (
	//       │  value = Name (
	//       │  id = 'self'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  attr = 'get_dnv_min_section_modulus'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Attribute
	//       │  args =  [
	//       │  Name (
	//       │  id = 'pressure'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  ] //
	//       │  keywords =  [
	//       │
	//       │  ] //
	//       │  starargs = None
	//       │  kwargs = None
	//       │  ) // Call
	//       │  ] //
	//       │  ) // Compare
	//       │  Error in func transpileOp
	//       │  └──not valid token: GtE for GtE (
	//       │
	//       │  ) // GtE
	//       └──Nil binary: Compare (
	//          left = Call (
	//          func = Name (
	//          id = 'min'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          args =  [
	//          Name (
	//          id = 'section_module'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          ] //
	//          keywords =  [
	//
	//          ] //
	//          starargs = None
	//          kwargs = None
	//          ) // Call
	//          ops =  [
	//          GtE (
	//
	//          ) // GtE
	//          ] //
	//          comparators =  [
	//          Call (
	//          func = Attribute (
	//          value = Name (
	//          id = 'self'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          attr = 'get_dnv_min_section_modulus'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Attribute
	//          args =  [
	//          Name (
	//          id = 'pressure'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          ] //
	//          keywords =  [
	//
	//          ] //
	//          starargs = None
	//          kwargs = None
	//          ) // Call
	//          ] //
	//          ) // Compare
	//
}
func is_acceptable_shear_area() {
	"\n        Returning if the shear area is ok.\n        :param shear_area:\n        :param pressure:\n        :return:\n        "
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in func transpileExprs
	//    │  ├──Error in Operation: Compare (
	//    │  │  left = Name (
	//    │  │  id = 'shear_area'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  ops =  [
	//    │  │  GtE (
	//    │  │
	//    │  │  ) // GtE
	//    │  │  ] //
	//    │  │  comparators =  [
	//    │  │  Call (
	//    │  │  func = Attribute (
	//    │  │  value = Name (
	//    │  │  id = 'self'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  attr = 'get_minimum_shear_area'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Attribute
	//    │  │  args =  [
	//    │  │  Name (
	//    │  │  id = 'pressure'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  ] //
	//    │  │  keywords =  [
	//    │  │
	//    │  │  ] //
	//    │  │  starargs = None
	//    │  │  kwargs = None
	//    │  │  ) // Call
	//    │  │  ] //
	//    │  │  ) // Compare
	//    │  │  Error in func transpileOp
	//    │  │  └──not valid token: GtE for GtE (
	//    │  │
	//    │  │  ) // GtE
	//    │  └──Nil binary: Compare (
	//    │     left = Name (
	//    │     id = 'shear_area'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     ops =  [
	//    │     GtE (
	//    │
	//    │     ) // GtE
	//    │     ] //
	//    │     comparators =  [
	//    │     Call (
	//    │     func = Attribute (
	//    │     value = Name (
	//    │     id = 'self'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     attr = 'get_minimum_shear_area'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Attribute
	//    │     args =  [
	//    │     Name (
	//    │     id = 'pressure'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     ] //
	//    │     keywords =  [
	//    │
	//    │     ] //
	//    │     starargs = None
	//    │     kwargs = None
	//    │     ) // Call
	//    │     ] //
	//    │     ) // Compare
	//    ├──Error function: 0
	//    └──Error in func transpileExprs
	//       ├──Error in Operation: Compare (
	//       │  left = Name (
	//       │  id = 'shear_area'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  ops =  [
	//       │  GtE (
	//       │
	//       │  ) // GtE
	//       │  ] //
	//       │  comparators =  [
	//       │  Call (
	//       │  func = Attribute (
	//       │  value = Name (
	//       │  id = 'self'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  attr = 'get_minimum_shear_area'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Attribute
	//       │  args =  [
	//       │  Name (
	//       │  id = 'pressure'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  ] //
	//       │  keywords =  [
	//       │
	//       │  ] //
	//       │  starargs = None
	//       │  kwargs = None
	//       │  ) // Call
	//       │  ] //
	//       │  ) // Compare
	//       │  Error in func transpileOp
	//       │  └──not valid token: GtE for GtE (
	//       │
	//       │  ) // GtE
	//       └──Nil binary: Compare (
	//          left = Name (
	//          id = 'shear_area'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          ops =  [
	//          GtE (
	//
	//          ) // GtE
	//          ] //
	//          comparators =  [
	//          Call (
	//          func = Attribute (
	//          value = Name (
	//          id = 'self'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          attr = 'get_minimum_shear_area'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Attribute
	//          args =  [
	//          Name (
	//          id = 'pressure'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          ] //
	//          keywords =  [
	//
	//          ] //
	//          starargs = None
	//          kwargs = None
	//          ) // Call
	//          ] //
	//          ) // Compare
	//
}
func get_plate_efficent_b() {
	"\n        Simple buckling calculations according to DNV-RP-C201\n        :return:\n        "
	s = self.spacing
	t = self.plate_th
	l = self.span
	E = 210000000000.0
	pSd = design_lat_press * 1000
	sigy1Sd = trans_stress_large * 1000000.0
	sigy2Sd = trans_stress_small * 1000000.0
	sigxSd = axial_stress * 1000000.0
	fy = self.mat_yield
	alphap = 0.525 * (s / t) * math.sqrt(fy/E)
	alphac = 1.1 * (s / t) * math.sqrt(fy/E)
	mu6_9 = 0.21 * (alphac - 0.2)
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in Operation: Compare (
	//    │  left = Name (
	//    │  id = 'alphac'
	//    │  ctx = Load (
	//    │
	//    │  ) // Load
	//    │  ) // Name
	//    │  ops =  [
	//    │  LtE (
	//    │
	//    │  ) // LtE
	//    │  ] //
	//    │  comparators =  [
	//    │  Num (
	//    │  n = 0.2
	//    │  ) // Num
	//    │  ] //
	//    │  ) // Compare
	//    │  Error in func transpileOp
	//    │  └──not valid token: LtE for LtE (
	//    │
	//    │  ) // LtE
	//    └──Nil binary: Compare (
	//       left = Name (
	//       id = 'alphac'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       LtE (
	//
	//       ) // LtE
	//       ] //
	//       comparators =  [
	//       Num (
	//       n = 0.2
	//       ) // Num
	//       ] //
	//       ) // Compare
	//
	ha = 0.05*(s/t) - 0.75
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'pSd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       LtE (
	//
	//       ) // LtE
	//       ] //
	//       comparators =  [
	//       BinOp (
	//       left = BinOp (
	//       left = Num (
	//       n = 2
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 't'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Pow (
	//
	//       ) // Pow
	//       right = Num (
	//       n = 2
	//       ) // Num
	//       ) // BinOp
	//       ) // BinOp
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ] //
	//       ) // Compare
	//       body = Num (
	//       n = 1
	//       ) // Num
	//       orelse = BinOp (
	//       left = Num (
	//       n = 1
	//       ) // Num
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Name (
	//       id = 'ha'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'pSd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Num (
	//       n = 2
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 't'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Pow (
	//
	//       ) // Pow
	//       right = Num (
	//       n = 2
	//       ) // Num
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       ) // IfExp
	//
	sigyR = (1.3*t/l*math.sqrt(E/fy) + kappa*(1-1.3*t/l*math.sqrt(E/fy))) * fy * kp
	l1 = min(0.25*l, 0.5*s)
	sig_min, sig_max = min(sigy1Sd, sigy2Sd), max(sigy1Sd, sigy2Sd)
	sigySd = sig_min + (sig_max-sig_min)*(1-l1/l)
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = BinOp (
	//       left = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 't'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ops =  [
	//       LtE (
	//
	//       ) // LtE
	//       ] //
	//       comparators =  [
	//       Num (
	//       n = 120
	//       ) // Num
	//       ] //
	//       ) // Compare
	//       body = BinOp (
	//       left = Num (
	//       n = 1
	//       ) // Num
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = BinOp (
	//       left = Num (
	//       n = 120
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 't'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       orelse = Num (
	//       n = 0
	//       ) // Num
	//       ) // IfExp
	//
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'alphap'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       Gt (
	//
	//       ) // Gt
	//       ] //
	//       comparators =  [
	//       Num (
	//       n = 0.673
	//       ) // Num
	//       ] //
	//       ) // Compare
	//       body = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'alphap'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Sub (
	//
	//       ) // Sub
	//       right = Num (
	//       n = 0.22
	//       ) // Num
	//       ) // BinOp
	//       op = Div (
	//
	//       ) // Div
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       Name (
	//       id = 'alphap'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       orelse = Num (
	//       n = 1
	//       ) // Num
	//       ) // IfExp
	//
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'sigySd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       GtE (
	//
	//       ) // GtE
	//       ] //
	//       comparators =  [
	//       Num (
	//       n = 0
	//       ) // Num
	//       ] //
	//       ) // Compare
	//       body = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'sqrt'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = BinOp (
	//       left = Num (
	//       n = 1
	//       ) // Num
	//       op = Sub (
	//
	//       ) // Sub
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Name (
	//       id = 'sigySd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'sigyR'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       op = Add (
	//
	//       ) // Add
	//       right = BinOp (
	//       left = Name (
	//       id = 'ci'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'sigxSd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'sigySd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Div (
	//
	//       ) // Div
	//       right = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'Cxs'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'sigyR'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       orelse = Call (
	//       func = Name (
	//       id = 'min'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       args =  [
	//       BinOp (
	//       left = Num (
	//       n = 0.5
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'sqrt'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Num (
	//       n = 4
	//       ) // Num
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Num (
	//       n = 3
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Name (
	//       id = 'sigySd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       ) // BinOp
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       op = Add (
	//
	//       ) // Add
	//       right = BinOp (
	//       left = Name (
	//       id = 'sigySd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       Num (
	//       n = 1
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // IfExp
	//
	return s * Cxs * Cys
}
func calculate_buckling_all() {
	"\n        Simple buckling calculations according to DNV-RP-C201\n        :return:\n        "
	As = self.web_height*self.web_th + self.flange_width*self.flange_th
	s = self.spacing
	t = self.plate_th
	l = self.span
	tf = self.flange_th
	tw = self.web_th
	hw = self.web_height
	bf = self.flange_width
	fy = self.mat_yield
	stf_type = self.get_stiffener_type()
	E = 210000000000.0
	Lg = 10
	mc = 13.3
	pSd = design_lat_press * 1000
	tauSd = self.tauxy * 1000000.0
	sigy1Sd = self.sigma_y1 * 1000000.0
	sigy2Sd = self.sigma_y2 * 1000000.0
	sigxSd = self.sigma_x * 1000000.0
	alphap = 0.525 * (s / t) * math.sqrt(fy/E)
	alphac = 1.1 * (s / t) * math.sqrt(fy/E)
	mu6_9 = 0.21 * (alphac - 0.2)
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in Operation: Compare (
	//    │  left = Name (
	//    │  id = 'alphac'
	//    │  ctx = Load (
	//    │
	//    │  ) // Load
	//    │  ) // Name
	//    │  ops =  [
	//    │  LtE (
	//    │
	//    │  ) // LtE
	//    │  ] //
	//    │  comparators =  [
	//    │  Num (
	//    │  n = 0.2
	//    │  ) // Num
	//    │  ] //
	//    │  ) // Compare
	//    │  Error in func transpileOp
	//    │  └──not valid token: LtE for LtE (
	//    │
	//    │  ) // LtE
	//    └──Nil binary: Compare (
	//       left = Name (
	//       id = 'alphac'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       LtE (
	//
	//       ) // LtE
	//       ] //
	//       comparators =  [
	//       Num (
	//       n = 0.2
	//       ) // Num
	//       ] //
	//       ) // Compare
	//
	ha = 0.05*(s/t) - 0.75
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'pSd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       LtE (
	//
	//       ) // LtE
	//       ] //
	//       comparators =  [
	//       BinOp (
	//       left = BinOp (
	//       left = Num (
	//       n = 2
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Name (
	//       id = 't'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ] //
	//       ) // Compare
	//       body = Num (
	//       n = 1
	//       ) // Num
	//       orelse = Call (
	//       func = Name (
	//       id = 'max'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       args =  [
	//       BinOp (
	//       left = Num (
	//       n = 1
	//       ) // Num
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Name (
	//       id = 'ha'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'pSd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Num (
	//       n = 2
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Name (
	//       id = 't'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       Num (
	//       n = 0
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // IfExp
	//
	sigyR = (1.3*t/l*math.sqrt(E/fy) + kappa*(1-1.3*t/l*math.sqrt(E/fy))) * fy * kp
	sigyRd = sigyR / 1.15
	ksp = math.sqrt(1 - 3*math.pow(tauSd/(fy/1), 2))
	l1 = min(0.25*l, 0.5*s)
	sig_min, sig_max = min(sigy1Sd, sigy2Sd), max(sigy1Sd, sigy2Sd)
	sigySd = sig_min + (sig_max-sig_min)*(1-l1/l)
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    └──Error in func transpileExprs
	//       ├──Error in Operation: Compare (
	//       │  left = Name (
	//       │  id = 'sigySd'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  ops =  [
	//       │  LtE (
	//       │
	//       │  ) // LtE
	//       │  ] //
	//       │  comparators =  [
	//       │  Name (
	//       │  id = 'sigyRd'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  ] //
	//       │  ) // Compare
	//       │  Error in func transpileOp
	//       │  └──not valid token: LtE for LtE (
	//       │
	//       │  ) // LtE
	//       └──Nil binary: Compare (
	//          left = Name (
	//          id = 'sigySd'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          ops =  [
	//          LtE (
	//
	//          ) // LtE
	//          ] //
	//          comparators =  [
	//          Name (
	//          id = 'sigyRd'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          ] //
	//          ) // Compare
	//
	Is = self.get_moment_of_intertia()
	Ip = math.pow(t, 3) * s / 10.9
	kc = 2 * (1 + math.sqrt(1+10.9*Is/(math.pow(t, 3)*s)))
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'l'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       LtE (
	//
	//       ) // LtE
	//       ] //
	//       comparators =  [
	//       Name (
	//       id = 'Lg'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ] //
	//       ) // Compare
	//       body = BinOp (
	//       left = Num (
	//       n = 5.34
	//       ) // Num
	//       op = Add (
	//
	//       ) // Add
	//       right = BinOp (
	//       left = Num (
	//       n = 4
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Name (
	//       id = 'l'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'Lg'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       ) // BinOp
	//       orelse = BinOp (
	//       left = BinOp (
	//       left = Num (
	//       n = 5.34
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Name (
	//       id = 'l'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'Lg'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       op = Add (
	//
	//       ) // Add
	//       right = Num (
	//       n = 4
	//       ) // Num
	//       ) // BinOp
	//       ) // IfExp
	//
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'l'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       GtE (
	//
	//       ) // GtE
	//       ] //
	//       comparators =  [
	//       Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ] //
	//       ) // Compare
	//       body = BinOp (
	//       left = Num (
	//       n = 5.34
	//       ) // Num
	//       op = Add (
	//
	//       ) // Add
	//       right = BinOp (
	//       left = Num (
	//       n = 4
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'l'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       ) // BinOp
	//       orelse = BinOp (
	//       left = BinOp (
	//       left = Num (
	//       n = 5.34
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'l'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       op = Add (
	//
	//       ) // Add
	//       right = Num (
	//       n = 4
	//       ) // Num
	//       ) // BinOp
	//       ) // IfExp
	//
	taucrg = kg * 0.904 * E * math.pow(t/l, 2)
	taucrl = kl * 0.904 * E * math.pow(t/s, 2)
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'tauSd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       Gt (
	//
	//       ) // Gt
	//       ] //
	//       comparators =  [
	//       BinOp (
	//       left = Name (
	//       id = 'taucrl'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Num (
	//       n = 1.15
	//       ) // Num
	//       ) // BinOp
	//       ] //
	//       ) // Compare
	//       body = BinOp (
	//       left = Name (
	//       id = 'tauSd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Sub (
	//
	//       ) // Sub
	//       right = Name (
	//       id = 'taucrg'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       orelse = Num (
	//       n = 0
	//       ) // Num
	//       ) // IfExp
	//
	taucrs = 36 * E / (s * t * math.pow(l, 2)) * math.Pow(Ip*math.pow(Is, 3), 0.25)
	tauRd = min(fy/(math.sqrt(3)*1.15), taucrl/1.15, taucrs/1.15)
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = BinOp (
	//       left = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 't'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ops =  [
	//       LtE (
	//
	//       ) // LtE
	//       ] //
	//       comparators =  [
	//       Num (
	//       n = 120
	//       ) // Num
	//       ] //
	//       ) // Compare
	//       body = BinOp (
	//       left = Num (
	//       n = 1
	//       ) // Num
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = BinOp (
	//       left = Num (
	//       n = 120
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 't'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       orelse = Num (
	//       n = 0
	//       ) // Num
	//       ) // IfExp
	//
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'alphap'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       Gt (
	//
	//       ) // Gt
	//       ] //
	//       comparators =  [
	//       Num (
	//       n = 0.673
	//       ) // Num
	//       ] //
	//       ) // Compare
	//       body = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'alphap'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Sub (
	//
	//       ) // Sub
	//       right = Num (
	//       n = 0.22
	//       ) // Num
	//       ) // BinOp
	//       op = Div (
	//
	//       ) // Div
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       Name (
	//       id = 'alphap'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       orelse = Num (
	//       n = 1
	//       ) // Num
	//       ) // IfExp
	//
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'sigySd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       GtE (
	//
	//       ) // GtE
	//       ] //
	//       comparators =  [
	//       Num (
	//       n = 0
	//       ) // Num
	//       ] //
	//       ) // Compare
	//       body = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'sqrt'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = BinOp (
	//       left = Num (
	//       n = 1
	//       ) // Num
	//       op = Sub (
	//
	//       ) // Sub
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Name (
	//       id = 'sigySd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'sigyR'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       op = Add (
	//
	//       ) // Add
	//       right = BinOp (
	//       left = Name (
	//       id = 'ci'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'sigxSd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'sigySd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Div (
	//
	//       ) // Div
	//       right = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'Cxs'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'sigyR'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       orelse = Call (
	//       func = Name (
	//       id = 'min'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       args =  [
	//       BinOp (
	//       left = Num (
	//       n = 0.5
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'sqrt'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Num (
	//       n = 4
	//       ) // Num
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Num (
	//       n = 3
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Call (
	//       func = Attribute (
	//       value = Name (
	//       id = 'math'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'pow'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       args =  [
	//       BinOp (
	//       left = Name (
	//       id = 'sigySd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       Num (
	//       n = 2
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // BinOp
	//       ) // BinOp
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       op = Add (
	//
	//       ) // Add
	//       right = BinOp (
	//       left = Name (
	//       id = 'sigySd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       Num (
	//       n = 1
	//       ) // Num
	//       ] //
	//       keywords =  [
	//
	//       ] //
	//       starargs = None
	//       kwargs = None
	//       ) // Call
	//       ) // IfExp
	//
	se = s * Cxs * Cys
	zp = self.get_cross_section_centroid_with_effective_plate(se) - t/2
	zt = t/2 + hw + tf - zp
	Ie = self.get_moment_of_intertia()
	Wep = Ie / zp
	Wes = Ie / zt
	C0 = Wes * fy * mc / (kc * E * math.pow(t, 2) * s)
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'psi'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       Gt (
	//
	//       ) // Gt
	//       ] //
	//       comparators =  [
	//       Num (
	//       n = 1.5
	//       ) // Num
	//       ] //
	//       ) // Compare
	//       body = BinOp (
	//       left = BinOp (
	//       left = BinOp (
	//       left = Num (
	//       n = 0.6
	//       ) // Num
	//       op = Add (
	//
	//       ) // Add
	//       right = BinOp (
	//       left = Num (
	//       n = 0.4
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'psi'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ) // BinOp
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'C0'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'sigy1Sd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       orelse = Num (
	//       n = 0
	//       ) // Num
	//       ) // IfExp
	//
	qSd = (pSd + p0) * s
	Ae = As + se*t
	W = min(Wes, Wep)
	pf = 12 * W / (math.pow(l, 2) * s) * (fy / 1.15)
	lk = l * (1 - 0.5*abs(pSd/pf))
	ie = math.sqrt(Ie / Ae)
	fE = math.pow(math.pi, 2) * E * math.pow(ie/lk, 2)
	sigjSD = math.sqrt(math.pow(sigxSd, 2) + math.pow(sigySd, 2) - sigxSd*sigySd + 3*math.pow(tauSd, 2))
	fEpx = 3.62 * E * math.pow(t/s, 2)
	fEpy = 0.9 * E * math.pow(t/s, 2)
	fEpt = 5.0 * E * math.pow(t/s, 2)
	c = 2 - s/l
	alphae = math.sqrt(fy / sigjSD * math.pow(math.pow(sigxSd/fEpx, c)+math.pow(sigySd/fEpy, c)+math.pow(tauSd/fEpt, c), 1/c))
	fep = fy / math.sqrt(1+math.pow(alphae, 4))
	eta = min(sigjSD/fep, 1)
	C = hw / s * math.pow(t/tw, 3) * math.sqrt(1-eta)
	beta = (3*C + 0.2) / (C + 0.2)
	Af = self.flange_width * self.flange_th
	Aw = self.web_height * self.web_th
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'stf_type'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       In (
	//
	//       ) // In
	//       ] //
	//       comparators =  [
	//       List (
	//       elts =  [
	//       Str (
	//       s = 'FB'
	//       ) // Str
	//       Str (
	//       s = 'T'
	//       ) // Str
	//       ] //
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // List
	//       ] //
	//       ) // Compare
	//       body = Num (
	//       n = 0
	//       ) // Num
	//       orelse = BinOp (
	//       left = BinOp (
	//       left = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'flange_width'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       op = Div (
	//
	//       ) // Div
	//       right = Num (
	//       n = 2
	//       ) // Num
	//       ) // BinOp
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Attribute (
	//       value = Name (
	//       id = 'self'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       attr = 'web_th'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Attribute
	//       op = Div (
	//
	//       ) // Div
	//       right = Num (
	//       n = 2
	//       ) // Num
	//       ) // BinOp
	//       ) // BinOp
	//       ) // IfExp
	//
	Iz = 1/12*Af*math.pow(bf, 2) + math.pow(ef, 2)*(Af/(1+Af/Aw))
	G = E / (2 * (1 + 0.3))
	lT = self.span
	u = math.pow(tauSd/tauRd, 2)
	fr, fks, fkp = get_some_data()
	Ms1Rd = Wes * (fr / 1.15)
	NksRd = Ae * (fks / 1.15)
	NkpRd = Ae * (fkp / 1.15)
	M1Sd = abs(qSd * math.pow(l, 2) / 12)
	M2Sd = abs(qSd * math.Pow(l, 2) / 24)
	Ne = math.pow(math.pi, 2) * E * Ae / math.pow(lk/ie, 2)
	Nrd = Ae * (fy / 1.15)
	Nsd = sigxSd*(As+s*t) + tautf*s*t
	zstar = 0
	MstRd = Wes * (fy / 1.15)
	MpRd = Wep * (fy / 1.15)
	fr, fks, fkp = get_some_data()
	Ms2Rd = Wes * (fr / 1.15)
	eq7_19 = sigySd / (ksp * sigyRd)
	if checked_side == "p" {
		// Find PY4GO error
		// Error in func transpileStmt
		// └──Error in func transpileExprs
		//    ├──Error function: 5
		//    └──Name: List
		//       cannot transpile : List (
		//       elts =  [
		//
		//       ] //
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // List
		//
		// Find PY4GO error
		// Error in func transpileStmt
		// └──Error in func transpileExprs
		//    ├──Error function: 5
		//    └──Name: List
		//       cannot transpile : List (
		//       elts =  [
		//
		//       ] //
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // List
		//
		for zstar := range np.arange(-zt/2, zp, 0.002) {
			eq7_50 = Nsd/NksRd + (M1Sd-Nsd*zstar)/(Ms1Rd*(1-Nsd/Ne)) + u
			eq7_51 = Nsd/NkpRd - 2*(Nsd/Nrd) + (M1Sd-Nsd*zstar)/(MpRd*(1-Nsd/Ne)) + u
			eq7_52 = Nsd/NksRd - 2*(Nsd/Nrd) + (M2Sd+Nsd*zstar)/(MstRd*(1-Nsd/Ne)) + u
			eq7_53 = Nsd/NkpRd + (M2Sd+Nsd*zstar)/(MpRd*(1-Nsd/Ne))
			max_lfs.append(max(eq7_50, eq7_51, eq7_52, eq7_53))
			// Find PY4GO error
			// Error in func transpileStmt
			// └──Error in func transpileExprs
			//    ├──Error in func transpileExprs
			//    │  └──Error in func transpileExprs
			//    │     └──Error in func transpileExprs
			//    │        ├──Error function: 5
			//    │        └──Name: List
			//    │           cannot transpile : List (
			//    │           elts =  [
			//    │           Name (
			//    │           id = 'eq7_19'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           Name (
			//    │           id = 'eq7_50'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           Name (
			//    │           id = 'eq7_51'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           Name (
			//    │           id = 'eq7_52'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           Name (
			//    │           id = 'eq7_53'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           Name (
			//    │           id = 'zstar'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           ] //
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // List
			//    ├──Error function: 0
			//    └──Error in func transpileExprs
			//       └──Error in func transpileExprs
			//          └──Error in func transpileExprs
			//             ├──Error function: 5
			//             └──Name: List
			//                cannot transpile : List (
			//                elts =  [
			//                Name (
			//                id = 'eq7_19'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                Name (
			//                id = 'eq7_50'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                Name (
			//                id = 'eq7_51'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                Name (
			//                id = 'eq7_52'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                Name (
			//                id = 'eq7_53'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                Name (
			//                id = 'zstar'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                ] //
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // List
			//
		}
		min_of_max_ufs_idx = max_lfs.index(min(max_lfs))
		return ufs[min_of_max_ufs_idx]
	} else {
		// Find PY4GO error
		// Error in func transpileStmt
		// └──Error in func transpileExprs
		//    ├──Error function: 5
		//    └──Name: List
		//       cannot transpile : List (
		//       elts =  [
		//
		//       ] //
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // List
		//
		// Find PY4GO error
		// Error in func transpileStmt
		// └──Error in func transpileExprs
		//    ├──Error function: 5
		//    └──Name: List
		//       cannot transpile : List (
		//       elts =  [
		//
		//       ] //
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // List
		//
		for zstar := range np.arange(-zt/2, zp, 0.002) {
			eq7_54 = Nsd/NksRd - 2*(Nsd/Nrd) + (M1Sd+Nsd*zstar)/(MstRd*(1-Nsd/Ne)) + u
			eq7_55 = Nsd/NkpRd + (M1Sd+Nsd*zstar)/(MpRd*(1-Nsd/Ne)) + u
			eq7_56 = Nsd/NksRd + (M2Sd-Nsd*zstar)/(Ms2Rd*(1-Nsd/Ne)) + u
			eq7_57 = Nsd/NkpRd - 2*(Nsd/Nrd) + (M2Sd-Nsd*zstar)/(MpRd*(1-Nsd/Ne)) + u
			max_lfs.append(max(eq7_54, eq7_55, eq7_56, eq7_57))
			// Find PY4GO error
			// Error in func transpileStmt
			// └──Error in func transpileExprs
			//    ├──Error in func transpileExprs
			//    │  └──Error in func transpileExprs
			//    │     └──Error in func transpileExprs
			//    │        ├──Error function: 5
			//    │        └──Name: List
			//    │           cannot transpile : List (
			//    │           elts =  [
			//    │           Name (
			//    │           id = 'eq7_19'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           Name (
			//    │           id = 'eq7_54'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           Name (
			//    │           id = 'eq7_55'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           Name (
			//    │           id = 'eq7_56'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           Name (
			//    │           id = 'eq7_57'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           Name (
			//    │           id = 'zstar'
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // Name
			//    │           ] //
			//    │           ctx = Load (
			//    │
			//    │           ) // Load
			//    │           ) // List
			//    ├──Error function: 0
			//    └──Error in func transpileExprs
			//       └──Error in func transpileExprs
			//          └──Error in func transpileExprs
			//             ├──Error function: 5
			//             └──Name: List
			//                cannot transpile : List (
			//                elts =  [
			//                Name (
			//                id = 'eq7_19'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                Name (
			//                id = 'eq7_54'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                Name (
			//                id = 'eq7_55'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                Name (
			//                id = 'eq7_56'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                Name (
			//                id = 'eq7_57'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                Name (
			//                id = 'zstar'
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // Name
			//                ] //
			//                ctx = Load (
			//
			//                ) // Load
			//                ) // List
			//
		}
		min_of_max_ufs_idx = max_lfs.index(min(max_lfs))
		return ufs[min_of_max_ufs_idx]
	}
}
func calculate_buckling_plate() {
	"\n        Simple buckling calculations according to DNV-RP-C201\n        This method is currently not used.\n        :return:\n        "
	s = self.spacing
	t = self.plate_th
	l = self.span
	E = 210000000000.0
	pSd = design_lat_press * 1000
	tauSd = design_shear_stress * 1000000.0
	sigy2Sd = trans_stress_small * 1000000.0
	fy = self.mat_yield
	alphac = 1.1 * (s / t) * math.sqrt(fy/E)
	gamma = 0.21 * (alphac - 0.2)
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in Operation: Compare (
	//    │  left = Name (
	//    │  id = 'alphac'
	//    │  ctx = Load (
	//    │
	//    │  ) // Load
	//    │  ) // Name
	//    │  ops =  [
	//    │  LtE (
	//    │
	//    │  ) // LtE
	//    │  ] //
	//    │  comparators =  [
	//    │  Num (
	//    │  n = 0.2
	//    │  ) // Num
	//    │  ] //
	//    │  ) // Compare
	//    │  Error in func transpileOp
	//    │  └──not valid token: LtE for LtE (
	//    │
	//    │  ) // LtE
	//    └──Nil binary: Compare (
	//       left = Name (
	//       id = 'alphac'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       LtE (
	//
	//       ) // LtE
	//       ] //
	//       comparators =  [
	//       Num (
	//       n = 0.2
	//       ) // Num
	//       ] //
	//       ) // Compare
	//
	ha = 0.05*(s/t) - 0.75
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: IfExp
	//       cannot transpile : IfExp (
	//       test = Compare (
	//       left = Name (
	//       id = 'pSd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       LtE (
	//
	//       ) // LtE
	//       ] //
	//       comparators =  [
	//       BinOp (
	//       left = BinOp (
	//       left = Num (
	//       n = 2
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 't'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Pow (
	//
	//       ) // Pow
	//       right = Num (
	//       n = 2
	//       ) // Num
	//       ) // BinOp
	//       ) // BinOp
	//       op = Mult (
	//
	//       ) // Mult
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       ] //
	//       ) // Compare
	//       body = Num (
	//       n = 1
	//       ) // Num
	//       orelse = BinOp (
	//       left = Num (
	//       n = 1
	//       ) // Num
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Name (
	//       id = 'ha'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 'pSd'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 'fy'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Sub (
	//
	//       ) // Sub
	//       right = BinOp (
	//       left = Num (
	//       n = 2
	//       ) // Num
	//       op = Mult (
	//
	//       ) // Mult
	//       right = BinOp (
	//       left = BinOp (
	//       left = Name (
	//       id = 't'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       op = Div (
	//
	//       ) // Div
	//       right = Name (
	//       id = 's'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ) // BinOp
	//       op = Pow (
	//
	//       ) // Pow
	//       right = Num (
	//       n = 2
	//       ) // Num
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       ) // BinOp
	//       ) // IfExp
	//
	sigyR = (1.3*t/l*math.sqrt(E/fy) + kappa*(1-1.3*t/l*math.sqrt(E/fy))) * fy * kp
	sigyRd = sigyR / 1.15
	ksp = math.sqrt(1 - 3*math.Pow(tauSd/(fy/1), 2))
	eq7_19 = ksp * sigyRd / sigy2Sd
	return eq7_19
}
func buckling_local_stiffener() {
	"\n        Local requirements for stiffeners. Chapter 9.11.\n        :return:\n        "
	epsilon = math.sqrt(235 / (self.mat_yield / 1000000.0))
	if self.stiffener_type == "L" {
		c = self.flange_width - self.web_th/2
	} else {
		if self.stiffener_type == "T" {
			c = self.flange_width/2 - self.web_th/2
		} else {
			if self.stiffener_type == "FB" {
				// Find PY4GO error
				// Error in func transpileStmt
				// └──Error in func transpileExprs
				//    ├──Error in func transpileExprs
				//    │  ├──Error function: 3
				//    │  └──Error in func transpileExprs
				//    │     ├──Error in Operation: Compare (
				//    │     │  left = Attribute (
				//    │     │  value = Name (
				//    │     │  id = 'self'
				//    │     │  ctx = Load (
				//    │     │
				//    │     │  ) // Load
				//    │     │  ) // Name
				//    │     │  attr = 'web_height'
				//    │     │  ctx = Load (
				//    │     │
				//    │     │  ) // Load
				//    │     │  ) // Attribute
				//    │     │  ops =  [
				//    │     │  LtE (
				//    │     │
				//    │     │  ) // LtE
				//    │     │  ] //
				//    │     │  comparators =  [
				//    │     │  BinOp (
				//    │     │  left = BinOp (
				//    │     │  left = Num (
				//    │     │  n = 42
				//    │     │  ) // Num
				//    │     │  op = Mult (
				//    │     │
				//    │     │  ) // Mult
				//    │     │  right = Attribute (
				//    │     │  value = Name (
				//    │     │  id = 'self'
				//    │     │  ctx = Load (
				//    │     │
				//    │     │  ) // Load
				//    │     │  ) // Name
				//    │     │  attr = 'web_th'
				//    │     │  ctx = Load (
				//    │     │
				//    │     │  ) // Load
				//    │     │  ) // Attribute
				//    │     │  ) // BinOp
				//    │     │  op = Mult (
				//    │     │
				//    │     │  ) // Mult
				//    │     │  right = Name (
				//    │     │  id = 'epsilon'
				//    │     │  ctx = Load (
				//    │     │
				//    │     │  ) // Load
				//    │     │  ) // Name
				//    │     │  ) // BinOp
				//    │     │  ] //
				//    │     │  ) // Compare
				//    │     │  Error in func transpileOp
				//    │     │  └──not valid token: LtE for LtE (
				//    │     │
				//    │     │  ) // LtE
				//    │     └──Nil binary: Compare (
				//    │        left = Attribute (
				//    │        value = Name (
				//    │        id = 'self'
				//    │        ctx = Load (
				//    │
				//    │        ) // Load
				//    │        ) // Name
				//    │        attr = 'web_height'
				//    │        ctx = Load (
				//    │
				//    │        ) // Load
				//    │        ) // Attribute
				//    │        ops =  [
				//    │        LtE (
				//    │
				//    │        ) // LtE
				//    │        ] //
				//    │        comparators =  [
				//    │        BinOp (
				//    │        left = BinOp (
				//    │        left = Num (
				//    │        n = 42
				//    │        ) // Num
				//    │        op = Mult (
				//    │
				//    │        ) // Mult
				//    │        right = Attribute (
				//    │        value = Name (
				//    │        id = 'self'
				//    │        ctx = Load (
				//    │
				//    │        ) // Load
				//    │        ) // Name
				//    │        attr = 'web_th'
				//    │        ctx = Load (
				//    │
				//    │        ) // Load
				//    │        ) // Attribute
				//    │        ) // BinOp
				//    │        op = Mult (
				//    │
				//    │        ) // Mult
				//    │        right = Name (
				//    │        id = 'epsilon'
				//    │        ctx = Load (
				//    │
				//    │        ) // Load
				//    │        ) // Name
				//    │        ) // BinOp
				//    │        ] //
				//    │        ) // Compare
				//    ├──Error function: 0
				//    └──Error in func transpileExprs
				//       ├──Error function: 3
				//       └──Error in func transpileExprs
				//          ├──Error in Operation: Compare (
				//          │  left = Attribute (
				//          │  value = Name (
				//          │  id = 'self'
				//          │  ctx = Load (
				//          │
				//          │  ) // Load
				//          │  ) // Name
				//          │  attr = 'web_height'
				//          │  ctx = Load (
				//          │
				//          │  ) // Load
				//          │  ) // Attribute
				//          │  ops =  [
				//          │  LtE (
				//          │
				//          │  ) // LtE
				//          │  ] //
				//          │  comparators =  [
				//          │  BinOp (
				//          │  left = BinOp (
				//          │  left = Num (
				//          │  n = 42
				//          │  ) // Num
				//          │  op = Mult (
				//          │
				//          │  ) // Mult
				//          │  right = Attribute (
				//          │  value = Name (
				//          │  id = 'self'
				//          │  ctx = Load (
				//          │
				//          │  ) // Load
				//          │  ) // Name
				//          │  attr = 'web_th'
				//          │  ctx = Load (
				//          │
				//          │  ) // Load
				//          │  ) // Attribute
				//          │  ) // BinOp
				//          │  op = Mult (
				//          │
				//          │  ) // Mult
				//          │  right = Name (
				//          │  id = 'epsilon'
				//          │  ctx = Load (
				//          │
				//          │  ) // Load
				//          │  ) // Name
				//          │  ) // BinOp
				//          │  ] //
				//          │  ) // Compare
				//          │  Error in func transpileOp
				//          │  └──not valid token: LtE for LtE (
				//          │
				//          │  ) // LtE
				//          └──Nil binary: Compare (
				//             left = Attribute (
				//             value = Name (
				//             id = 'self'
				//             ctx = Load (
				//
				//             ) // Load
				//             ) // Name
				//             attr = 'web_height'
				//             ctx = Load (
				//
				//             ) // Load
				//             ) // Attribute
				//             ops =  [
				//             LtE (
				//
				//             ) // LtE
				//             ] //
				//             comparators =  [
				//             BinOp (
				//             left = BinOp (
				//             left = Num (
				//             n = 42
				//             ) // Num
				//             op = Mult (
				//
				//             ) // Mult
				//             right = Attribute (
				//             value = Name (
				//             id = 'self'
				//             ctx = Load (
				//
				//             ) // Load
				//             ) // Name
				//             attr = 'web_th'
				//             ctx = Load (
				//
				//             ) // Load
				//             ) // Attribute
				//             ) // BinOp
				//             op = Mult (
				//
				//             ) // Mult
				//             right = Name (
				//             id = 'epsilon'
				//             ctx = Load (
				//
				//             ) // Load
				//             ) // Name
				//             ) // BinOp
				//             ] //
				//             ) // Compare
				//
			}
		}
	}
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in func transpileExprs
	//    │  ├──Error function: 3
	//    │  └──Error in func transpileExprs
	//    │     ├──Error in func transpileOp
	//    │     │  └──not valid token: And for And (
	//    │     │
	//    │     │     ) // And
	//    │     ├──Error in func transpileExprs
	//    │     │  ├──Error in func transpileExprs
	//    │     │  │  ├──Error in Operation: Compare (
	//    │     │  │  │  left = Name (
	//    │     │  │  │  id = 'c'
	//    │     │  │  │  ctx = Load (
	//    │     │  │  │
	//    │     │  │  │  ) // Load
	//    │     │  │  │  ) // Name
	//    │     │  │  │  ops =  [
	//    │     │  │  │  LtE (
	//    │     │  │  │
	//    │     │  │  │  ) // LtE
	//    │     │  │  │  ] //
	//    │     │  │  │  comparators =  [
	//    │     │  │  │  BinOp (
	//    │     │  │  │  left = BinOp (
	//    │     │  │  │  left = Num (
	//    │     │  │  │  n = 14
	//    │     │  │  │  ) // Num
	//    │     │  │  │  op = Mult (
	//    │     │  │  │
	//    │     │  │  │  ) // Mult
	//    │     │  │  │  right = Attribute (
	//    │     │  │  │  value = Name (
	//    │     │  │  │  id = 'self'
	//    │     │  │  │  ctx = Load (
	//    │     │  │  │
	//    │     │  │  │  ) // Load
	//    │     │  │  │  ) // Name
	//    │     │  │  │  attr = 'flange_th'
	//    │     │  │  │  ctx = Load (
	//    │     │  │  │
	//    │     │  │  │  ) // Load
	//    │     │  │  │  ) // Attribute
	//    │     │  │  │  ) // BinOp
	//    │     │  │  │  op = Mult (
	//    │     │  │  │
	//    │     │  │  │  ) // Mult
	//    │     │  │  │  right = Name (
	//    │     │  │  │  id = 'epsilon'
	//    │     │  │  │  ctx = Load (
	//    │     │  │  │
	//    │     │  │  │  ) // Load
	//    │     │  │  │  ) // Name
	//    │     │  │  │  ) // BinOp
	//    │     │  │  │  ] //
	//    │     │  │  │  ) // Compare
	//    │     │  │  │  Error in func transpileOp
	//    │     │  │  │  └──not valid token: LtE for LtE (
	//    │     │  │  │
	//    │     │  │  │  ) // LtE
	//    │     │  │  └──Nil binary: Compare (
	//    │     │  │     left = Name (
	//    │     │  │     id = 'c'
	//    │     │  │     ctx = Load (
	//    │     │  │
	//    │     │  │     ) // Load
	//    │     │  │     ) // Name
	//    │     │  │     ops =  [
	//    │     │  │     LtE (
	//    │     │  │
	//    │     │  │     ) // LtE
	//    │     │  │     ] //
	//    │     │  │     comparators =  [
	//    │     │  │     BinOp (
	//    │     │  │     left = BinOp (
	//    │     │  │     left = Num (
	//    │     │  │     n = 14
	//    │     │  │     ) // Num
	//    │     │  │     op = Mult (
	//    │     │  │
	//    │     │  │     ) // Mult
	//    │     │  │     right = Attribute (
	//    │     │  │     value = Name (
	//    │     │  │     id = 'self'
	//    │     │  │     ctx = Load (
	//    │     │  │
	//    │     │  │     ) // Load
	//    │     │  │     ) // Name
	//    │     │  │     attr = 'flange_th'
	//    │     │  │     ctx = Load (
	//    │     │  │
	//    │     │  │     ) // Load
	//    │     │  │     ) // Attribute
	//    │     │  │     ) // BinOp
	//    │     │  │     op = Mult (
	//    │     │  │
	//    │     │  │     ) // Mult
	//    │     │  │     right = Name (
	//    │     │  │     id = 'epsilon'
	//    │     │  │     ctx = Load (
	//    │     │  │
	//    │     │  │     ) // Load
	//    │     │  │     ) // Name
	//    │     │  │     ) // BinOp
	//    │     │  │     ] //
	//    │     │  │     ) // Compare
	//    │     │  └──Error in func transpileExprs
	//    │     │     ├──Error in Operation: Compare (
	//    │     │     │  left = Attribute (
	//    │     │     │  value = Name (
	//    │     │     │  id = 'self'
	//    │     │     │  ctx = Load (
	//    │     │     │
	//    │     │     │  ) // Load
	//    │     │     │  ) // Name
	//    │     │     │  attr = 'web_height'
	//    │     │     │  ctx = Load (
	//    │     │     │
	//    │     │     │  ) // Load
	//    │     │     │  ) // Attribute
	//    │     │     │  ops =  [
	//    │     │     │  LtE (
	//    │     │     │
	//    │     │     │  ) // LtE
	//    │     │     │  ] //
	//    │     │     │  comparators =  [
	//    │     │     │  BinOp (
	//    │     │     │  left = BinOp (
	//    │     │     │  left = Num (
	//    │     │     │  n = 42
	//    │     │     │  ) // Num
	//    │     │     │  op = Mult (
	//    │     │     │
	//    │     │     │  ) // Mult
	//    │     │     │  right = Attribute (
	//    │     │     │  value = Name (
	//    │     │     │  id = 'self'
	//    │     │     │  ctx = Load (
	//    │     │     │
	//    │     │     │  ) // Load
	//    │     │     │  ) // Name
	//    │     │     │  attr = 'web_th'
	//    │     │     │  ctx = Load (
	//    │     │     │
	//    │     │     │  ) // Load
	//    │     │     │  ) // Attribute
	//    │     │     │  ) // BinOp
	//    │     │     │  op = Mult (
	//    │     │     │
	//    │     │     │  ) // Mult
	//    │     │     │  right = Name (
	//    │     │     │  id = 'epsilon'
	//    │     │     │  ctx = Load (
	//    │     │     │
	//    │     │     │  ) // Load
	//    │     │     │  ) // Name
	//    │     │     │  ) // BinOp
	//    │     │     │  ] //
	//    │     │     │  ) // Compare
	//    │     │     │  Error in func transpileOp
	//    │     │     │  └──not valid token: LtE for LtE (
	//    │     │     │
	//    │     │     │  ) // LtE
	//    │     │     └──Nil binary: Compare (
	//    │     │        left = Attribute (
	//    │     │        value = Name (
	//    │     │        id = 'self'
	//    │     │        ctx = Load (
	//    │     │
	//    │     │        ) // Load
	//    │     │        ) // Name
	//    │     │        attr = 'web_height'
	//    │     │        ctx = Load (
	//    │     │
	//    │     │        ) // Load
	//    │     │        ) // Attribute
	//    │     │        ops =  [
	//    │     │        LtE (
	//    │     │
	//    │     │        ) // LtE
	//    │     │        ] //
	//    │     │        comparators =  [
	//    │     │        BinOp (
	//    │     │        left = BinOp (
	//    │     │        left = Num (
	//    │     │        n = 42
	//    │     │        ) // Num
	//    │     │        op = Mult (
	//    │     │
	//    │     │        ) // Mult
	//    │     │        right = Attribute (
	//    │     │        value = Name (
	//    │     │        id = 'self'
	//    │     │        ctx = Load (
	//    │     │
	//    │     │        ) // Load
	//    │     │        ) // Name
	//    │     │        attr = 'web_th'
	//    │     │        ctx = Load (
	//    │     │
	//    │     │        ) // Load
	//    │     │        ) // Attribute
	//    │     │        ) // BinOp
	//    │     │        op = Mult (
	//    │     │
	//    │     │        ) // Mult
	//    │     │        right = Name (
	//    │     │        id = 'epsilon'
	//    │     │        ctx = Load (
	//    │     │
	//    │     │        ) // Load
	//    │     │        ) // Name
	//    │     │        ) // BinOp
	//    │     │        ] //
	//    │     │        ) // Compare
	//    │     └──Compare is not enought
	//    ├──Error function: 0
	//    └──Error in func transpileExprs
	//       ├──Error function: 3
	//       └──Error in func transpileExprs
	//          ├──Error in func transpileOp
	//          │  └──not valid token: And for And (
	//          │
	//          │     ) // And
	//          ├──Error in func transpileExprs
	//          │  ├──Error in func transpileExprs
	//          │  │  ├──Error in Operation: Compare (
	//          │  │  │  left = Name (
	//          │  │  │  id = 'c'
	//          │  │  │  ctx = Load (
	//          │  │  │
	//          │  │  │  ) // Load
	//          │  │  │  ) // Name
	//          │  │  │  ops =  [
	//          │  │  │  LtE (
	//          │  │  │
	//          │  │  │  ) // LtE
	//          │  │  │  ] //
	//          │  │  │  comparators =  [
	//          │  │  │  BinOp (
	//          │  │  │  left = BinOp (
	//          │  │  │  left = Num (
	//          │  │  │  n = 14
	//          │  │  │  ) // Num
	//          │  │  │  op = Mult (
	//          │  │  │
	//          │  │  │  ) // Mult
	//          │  │  │  right = Attribute (
	//          │  │  │  value = Name (
	//          │  │  │  id = 'self'
	//          │  │  │  ctx = Load (
	//          │  │  │
	//          │  │  │  ) // Load
	//          │  │  │  ) // Name
	//          │  │  │  attr = 'flange_th'
	//          │  │  │  ctx = Load (
	//          │  │  │
	//          │  │  │  ) // Load
	//          │  │  │  ) // Attribute
	//          │  │  │  ) // BinOp
	//          │  │  │  op = Mult (
	//          │  │  │
	//          │  │  │  ) // Mult
	//          │  │  │  right = Name (
	//          │  │  │  id = 'epsilon'
	//          │  │  │  ctx = Load (
	//          │  │  │
	//          │  │  │  ) // Load
	//          │  │  │  ) // Name
	//          │  │  │  ) // BinOp
	//          │  │  │  ] //
	//          │  │  │  ) // Compare
	//          │  │  │  Error in func transpileOp
	//          │  │  │  └──not valid token: LtE for LtE (
	//          │  │  │
	//          │  │  │  ) // LtE
	//          │  │  └──Nil binary: Compare (
	//          │  │     left = Name (
	//          │  │     id = 'c'
	//          │  │     ctx = Load (
	//          │  │
	//          │  │     ) // Load
	//          │  │     ) // Name
	//          │  │     ops =  [
	//          │  │     LtE (
	//          │  │
	//          │  │     ) // LtE
	//          │  │     ] //
	//          │  │     comparators =  [
	//          │  │     BinOp (
	//          │  │     left = BinOp (
	//          │  │     left = Num (
	//          │  │     n = 14
	//          │  │     ) // Num
	//          │  │     op = Mult (
	//          │  │
	//          │  │     ) // Mult
	//          │  │     right = Attribute (
	//          │  │     value = Name (
	//          │  │     id = 'self'
	//          │  │     ctx = Load (
	//          │  │
	//          │  │     ) // Load
	//          │  │     ) // Name
	//          │  │     attr = 'flange_th'
	//          │  │     ctx = Load (
	//          │  │
	//          │  │     ) // Load
	//          │  │     ) // Attribute
	//          │  │     ) // BinOp
	//          │  │     op = Mult (
	//          │  │
	//          │  │     ) // Mult
	//          │  │     right = Name (
	//          │  │     id = 'epsilon'
	//          │  │     ctx = Load (
	//          │  │
	//          │  │     ) // Load
	//          │  │     ) // Name
	//          │  │     ) // BinOp
	//          │  │     ] //
	//          │  │     ) // Compare
	//          │  └──Error in func transpileExprs
	//          │     ├──Error in Operation: Compare (
	//          │     │  left = Attribute (
	//          │     │  value = Name (
	//          │     │  id = 'self'
	//          │     │  ctx = Load (
	//          │     │
	//          │     │  ) // Load
	//          │     │  ) // Name
	//          │     │  attr = 'web_height'
	//          │     │  ctx = Load (
	//          │     │
	//          │     │  ) // Load
	//          │     │  ) // Attribute
	//          │     │  ops =  [
	//          │     │  LtE (
	//          │     │
	//          │     │  ) // LtE
	//          │     │  ] //
	//          │     │  comparators =  [
	//          │     │  BinOp (
	//          │     │  left = BinOp (
	//          │     │  left = Num (
	//          │     │  n = 42
	//          │     │  ) // Num
	//          │     │  op = Mult (
	//          │     │
	//          │     │  ) // Mult
	//          │     │  right = Attribute (
	//          │     │  value = Name (
	//          │     │  id = 'self'
	//          │     │  ctx = Load (
	//          │     │
	//          │     │  ) // Load
	//          │     │  ) // Name
	//          │     │  attr = 'web_th'
	//          │     │  ctx = Load (
	//          │     │
	//          │     │  ) // Load
	//          │     │  ) // Attribute
	//          │     │  ) // BinOp
	//          │     │  op = Mult (
	//          │     │
	//          │     │  ) // Mult
	//          │     │  right = Name (
	//          │     │  id = 'epsilon'
	//          │     │  ctx = Load (
	//          │     │
	//          │     │  ) // Load
	//          │     │  ) // Name
	//          │     │  ) // BinOp
	//          │     │  ] //
	//          │     │  ) // Compare
	//          │     │  Error in func transpileOp
	//          │     │  └──not valid token: LtE for LtE (
	//          │     │
	//          │     │  ) // LtE
	//          │     └──Nil binary: Compare (
	//          │        left = Attribute (
	//          │        value = Name (
	//          │        id = 'self'
	//          │        ctx = Load (
	//          │
	//          │        ) // Load
	//          │        ) // Name
	//          │        attr = 'web_height'
	//          │        ctx = Load (
	//          │
	//          │        ) // Load
	//          │        ) // Attribute
	//          │        ops =  [
	//          │        LtE (
	//          │
	//          │        ) // LtE
	//          │        ] //
	//          │        comparators =  [
	//          │        BinOp (
	//          │        left = BinOp (
	//          │        left = Num (
	//          │        n = 42
	//          │        ) // Num
	//          │        op = Mult (
	//          │
	//          │        ) // Mult
	//          │        right = Attribute (
	//          │        value = Name (
	//          │        id = 'self'
	//          │        ctx = Load (
	//          │
	//          │        ) // Load
	//          │        ) // Name
	//          │        attr = 'web_th'
	//          │        ctx = Load (
	//          │
	//          │        ) // Load
	//          │        ) // Attribute
	//          │        ) // BinOp
	//          │        op = Mult (
	//          │
	//          │        ) // Mult
	//          │        right = Name (
	//          │        id = 'epsilon'
	//          │        ctx = Load (
	//          │
	//          │        ) // Load
	//          │        ) // Name
	//          │        ) // BinOp
	//          │        ] //
	//          │        ) // Compare
	//          └──Compare is not enought
	//
}
func is_acceptable_pl_thk() {
	"\n        Checking if the thickness is acceptable.\n        :return:\n        "
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in func transpileExprs
	//    │  ├──Error in Operation: Compare (
	//    │  │  left = Call (
	//    │  │  func = Attribute (
	//    │  │  value = Name (
	//    │  │  id = 'self'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  attr = 'get_dnv_min_thickness'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Attribute
	//    │  │  args =  [
	//    │  │  Name (
	//    │  │  id = 'design_pressure'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  ] //
	//    │  │  keywords =  [
	//    │  │
	//    │  │  ] //
	//    │  │  starargs = None
	//    │  │  kwargs = None
	//    │  │  ) // Call
	//    │  │  ops =  [
	//    │  │  LtE (
	//    │  │
	//    │  │  ) // LtE
	//    │  │  ] //
	//    │  │  comparators =  [
	//    │  │  BinOp (
	//    │  │  left = Attribute (
	//    │  │  value = Name (
	//    │  │  id = 'self'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Name
	//    │  │  attr = 'plate_th'
	//    │  │  ctx = Load (
	//    │  │
	//    │  │  ) // Load
	//    │  │  ) // Attribute
	//    │  │  op = Mult (
	//    │  │
	//    │  │  ) // Mult
	//    │  │  right = Num (
	//    │  │  n = 1000
	//    │  │  ) // Num
	//    │  │  ) // BinOp
	//    │  │  ] //
	//    │  │  ) // Compare
	//    │  │  Error in func transpileOp
	//    │  │  └──not valid token: LtE for LtE (
	//    │  │
	//    │  │  ) // LtE
	//    │  └──Nil binary: Compare (
	//    │     left = Call (
	//    │     func = Attribute (
	//    │     value = Name (
	//    │     id = 'self'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     attr = 'get_dnv_min_thickness'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Attribute
	//    │     args =  [
	//    │     Name (
	//    │     id = 'design_pressure'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     ] //
	//    │     keywords =  [
	//    │
	//    │     ] //
	//    │     starargs = None
	//    │     kwargs = None
	//    │     ) // Call
	//    │     ops =  [
	//    │     LtE (
	//    │
	//    │     ) // LtE
	//    │     ] //
	//    │     comparators =  [
	//    │     BinOp (
	//    │     left = Attribute (
	//    │     value = Name (
	//    │     id = 'self'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Name
	//    │     attr = 'plate_th'
	//    │     ctx = Load (
	//    │
	//    │     ) // Load
	//    │     ) // Attribute
	//    │     op = Mult (
	//    │
	//    │     ) // Mult
	//    │     right = Num (
	//    │     n = 1000
	//    │     ) // Num
	//    │     ) // BinOp
	//    │     ] //
	//    │     ) // Compare
	//    ├──Error function: 0
	//    └──Error in func transpileExprs
	//       ├──Error in Operation: Compare (
	//       │  left = Call (
	//       │  func = Attribute (
	//       │  value = Name (
	//       │  id = 'self'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  attr = 'get_dnv_min_thickness'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Attribute
	//       │  args =  [
	//       │  Name (
	//       │  id = 'design_pressure'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  ] //
	//       │  keywords =  [
	//       │
	//       │  ] //
	//       │  starargs = None
	//       │  kwargs = None
	//       │  ) // Call
	//       │  ops =  [
	//       │  LtE (
	//       │
	//       │  ) // LtE
	//       │  ] //
	//       │  comparators =  [
	//       │  BinOp (
	//       │  left = Attribute (
	//       │  value = Name (
	//       │  id = 'self'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Name
	//       │  attr = 'plate_th'
	//       │  ctx = Load (
	//       │
	//       │  ) // Load
	//       │  ) // Attribute
	//       │  op = Mult (
	//       │
	//       │  ) // Mult
	//       │  right = Num (
	//       │  n = 1000
	//       │  ) // Num
	//       │  ) // BinOp
	//       │  ] //
	//       │  ) // Compare
	//       │  Error in func transpileOp
	//       │  └──not valid token: LtE for LtE (
	//       │
	//       │  ) // LtE
	//       └──Nil binary: Compare (
	//          left = Call (
	//          func = Attribute (
	//          value = Name (
	//          id = 'self'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          attr = 'get_dnv_min_thickness'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Attribute
	//          args =  [
	//          Name (
	//          id = 'design_pressure'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          ] //
	//          keywords =  [
	//
	//          ] //
	//          starargs = None
	//          kwargs = None
	//          ) // Call
	//          ops =  [
	//          LtE (
	//
	//          ) // LtE
	//          ] //
	//          comparators =  [
	//          BinOp (
	//          left = Attribute (
	//          value = Name (
	//          id = 'self'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Name
	//          attr = 'plate_th'
	//          ctx = Load (
	//
	//          ) // Load
	//          ) // Attribute
	//          op = Mult (
	//
	//          ) // Mult
	//          right = Num (
	//          n = 1000
	//          ) // Num
	//          ) // BinOp
	//          ] //
	//          ) // Compare
	//
}
func __init__() {
	super(CalcFatigue, self).__init__(main_dict, fatigue_dict)
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error in Operation: Compare (
	//    │  left = Name (
	//    │  id = 'fatigue_dict'
	//    │  ctx = Load (
	//    │
	//    │  ) // Load
	//    │  ) // Name
	//    │  ops =  [
	//    │  IsNot (
	//    │
	//    │  ) // IsNot
	//    │  ] //
	//    │  comparators =  [
	//    │  Name (
	//    │  id = 'None'
	//    │  ctx = Load (
	//    │
	//    │  ) // Load
	//    │  ) // Name
	//    │  ] //
	//    │  ) // Compare
	//    │  Error in func transpileOp
	//    │  └──not valid token: IsNot for IsNot (
	//    │
	//    │  ) // IsNot
	//    └──Nil binary: Compare (
	//       left = Name (
	//       id = 'fatigue_dict'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ops =  [
	//       IsNot (
	//
	//       ) // IsNot
	//       ] //
	//       comparators =  [
	//       Name (
	//       id = 'None'
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // Name
	//       ] //
	//       ) // Compare
	//
}
func get_sn_curve() {
	return self._sn_curve
}
func __get_sigma_ext() {
	return 0.5 * int_press * math.Pow(self.spacing/self.plate_th, 2) * (self._k_factor / math.Pow(1000, 2))
}
func __get_sigma_int() {
	return 0.5 * ext_press * math.Pow(self.spacing/self.plate_th, 2) * (self._k_factor / math.Pow(1000, 2))
}
func __get_range() {
	return 2 * math.sqrt(math.pow(self.__get_sigma_ext(ext_press), 2)+math.pow(self.__get_sigma_int(int_press), 2)+2*self._corr_loc[idx]*self.__get_sigma_ext(ext_press)*self.__get_sigma_int(int_press))
}
func __get_stress_fraction() {
	return self.__get_range(idx, int_press, ext_press) / math.pow(math.log(self._no_of_cycles), 1/self._weibull[idx])
}
func __get_gamma1() {
	return math.exp(gammaln(snc.get_paramter(self._sn_curve, "m1")/self._weibull[idx] + 1))
}
func __get_gamma2() {
	return math.exp(gammaln(snc.get_paramter(self._sn_curve, "m2")/self._weibull[idx] + 1))
}
func get_damage_slope1() {
	m1, log_a1, k, slope = snc.get_paramter(curve, "m1"), snc.get_paramter(curve, "log a1"), snc.get_paramter(curve, "k"), snc.get_paramter(curve, "slope")
	cycles = self._design_life * 365 * 24 * 3600 / self._period[idx]
	thk_eff = math.log10(max(1, self.plate_th/0.025)) * k
	slope_ch = math.exp(math.log(math.pow(10, log_a1-m1*thk_eff)/slope) / m1)
	gamma1 = self.__get_gamma1(idx)
	weibull = self._weibull[idx]
	stress_frac = self.__get_stress_fraction(idx, int_press, ext_press)
	if stress_frac == 0 {
		return 0
	}
	x, alpha = math.pow(slope_ch/stress_frac, weibull), 1+m1/weibull
	gamma_val = gammadist.cdf(x, alpha)
	return cycles / math.pow(10, log_a1-m1*thk_eff) * math.pow(stress_frac, m1) * gamma1 * (1 - gamma_val) * self._fraction[idx]
}
func get_damage_slope2() {
	m2, log_m2, k, slope = snc.get_paramter(curve, "m2"), snc.get_paramter(curve, "log a2"), snc.get_paramter(curve, "k"), snc.get_paramter(curve, "slope")
	cycles = self._design_life * 365 * 24 * 3600 / self._period[idx]
	thk_eff = math.log10(max(1, self.plate_th/25)) * k
	slope_ch = math.exp(math.log(math.pow(10, log_m2-m2*thk_eff)/slope) / m2)
	gammm2 = self.__get_gamma2(idx)
	weibull = self._weibull[idx]
	stress_frac = self.__get_stress_fraction(idx, int_press, ext_press)
	if stress_frac == 0 {
		return 0
	}
	x, alpha = math.pow(slope_ch/stress_frac, weibull), 1+m2/weibull
	gamma_val = gammadist.cdf(x, alpha)
	return cycles / math.pow(10, log_m2-m2*thk_eff) * math.pow(stress_frac, m2) * gammm2 * gamma_val * self._fraction[idx]
}
func get_total_damage() {
	damage = 0
	for idx := range range(3) {
		// Find PY4GO error
		// Error in func transpileStmt
		// └──Error in func transpileExprs
		//    └──Error in func transpileOp
		//       └──not valid token: And for And (
		//
		//          ) // And
		//
	}
	return damage
}
func set_fatigue_properties() {
	" Setting the fatiuge properties. "
	self._sn_curve, self.fatigue_dict["SN-curve"] = fatigue_dict["SN-curve"], fatigue_dict["SN-curve"]
	self._acc, self.fatigue_dict["Accelerations"] = fatigue_dict["Accelerations"], fatigue_dict["Accelerations"]
	self._weibull, self.fatigue_dict["Weibull"] = fatigue_dict["Weibull"], fatigue_dict["Weibull"]
	self._period, self.fatigue_dict["Period"] = fatigue_dict["Period"], fatigue_dict["Period"]
	self._k_factor, self.fatigue_dict["SCF"] = fatigue_dict["SCF"], fatigue_dict["SCF"]
	self._corr_loc, self.fatigue_dict["CorrLoc"] = fatigue_dict["CorrLoc"], fatigue_dict["CorrLoc"]
	self._no_of_cycles, self.fatigue_dict["n0"] = fatigue_dict["n0"], fatigue_dict["n0"]
	self._design_life, self.fatigue_dict["Design life"] = fatigue_dict["Design life"], fatigue_dict["Design life"]
	self._fraction, self.fatigue_dict["Fraction"] = fatigue_dict["Fraction"], fatigue_dict["Fraction"]
	self._case_order, self.fatigue_dict["Order"] = fatigue_dict["Order"], fatigue_dict["Order"]
	self._dff, self.fatigue_dict["DFF"] = fatigue_dict["DFF"], fatigue_dict["DFF"]
}
func get_fatigue_properties() {
	" Returning properties as a dictionary "
	return self.fatigue_dict
}
func get_accelerations() {
	" Returning tuple of accelerattions."
	return self._acc
}
func get_dff() {
	return self._dff
}
func get_design_life() {
	return self._design_life
}
func main() {
	"\n    Setting the properties for the plate and the stiffener. Takes a dictionary as argument.\n    "
	"\n    This Class does the calculations for the plate fields. \n    Input is a structure object, same as for the structure class.\n    The class inherits from Structure class.\n    "
	"\n    This Class does the calculations for the plate fields. \n    Input is a structure object (getters from the Structure Class)\n    "
	if __name__ == "__main__" {
		// Find PY4GO error
		// Error in func transpileStmt
		// └──Error in func transpileExprs
		//    ├──Error function: 5
		//    └──Name: List
		//       cannot transpile : List (
		//       elts =  [
		//       Call (
		//       func = Name (
		//       id = 'CalcScantlings'
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // Name
		//       args =  [
		//       Attribute (
		//       value = Name (
		//       id = 'ex'
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // Name
		//       attr = 'obj_dict'
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // Attribute
		//       ] //
		//       keywords =  [
		//
		//       ] //
		//       starargs = None
		//       kwargs = None
		//       ) // Call
		//       Call (
		//       func = Name (
		//       id = 'CalcScantlings'
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // Name
		//       args =  [
		//       Attribute (
		//       value = Name (
		//       id = 'ex'
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // Name
		//       attr = 'obj_dict2'
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // Attribute
		//       ] //
		//       keywords =  [
		//
		//       ] //
		//       starargs = None
		//       kwargs = None
		//       ) // Call
		//       Call (
		//       func = Name (
		//       id = 'CalcScantlings'
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // Name
		//       args =  [
		//       Attribute (
		//       value = Name (
		//       id = 'ex'
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // Name
		//       attr = 'obj_dict_L'
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // Attribute
		//       ] //
		//       keywords =  [
		//
		//       ] //
		//       starargs = None
		//       kwargs = None
		//       ) // Call
		//       ] //
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // List
		//
	}
}
