package main

func __init__() {
	self.main_load_dict = main_load_dict
	self.static_draft = main_load_dict["static_draft"]
	self.poly_third = main_load_dict["poly_third"]
	self.poly_second = main_load_dict["poly_second"]
	self.poly_first = main_load_dict["poly_first"]
	self.poly_const = main_load_dict["poly_const"]
	self.manual_press = main_load_dict["man_press"]
	self.load_condition = main_load_dict["load_condition"]
	self.name_of_load = main_load_dict["name_of_load"]
	self.dynamic_pressure = 0
	self.static_pressure = 0
	self.is_external = True
}
func __str__() {
	string = str("Properties selected load is:" + "\n----------------------------" + "\n Name of load: " + str(self.name_of_load) + "\n Polynominal (x^3): " + str(self.poly_third) + "\n Polynominal (x^2): " + str(self.poly_second) + "\n Polynominal (x):   " + str(self.poly_first) + "\n Constant (C):      " + str(self.poly_const) + "\n Load condition:    " + str(self.load_condition) + "\n Limit state       " + str(self.limit_state) + "\n Is external?       " + str(self.is_external) + "\n Static draft:      " + str(self.static_draft))
	return string
}
func get_calculated_pressure() {
	"\n        Input variable is a tuple of (x,y). This method need one variable and the right one must be chosen.\n        :param varibale_value:\n        :return:\n        "
	input_var = varibale_value
	if self.is_static() {
		press = 1025 * acceleration * (self.static_draft - input_var[1])
	} else {
		// Find PY4GO error
		// Error in func transpileStmt
		// └──Error in func transpileExprs
		//    ├──Error in Operation: Compare (
		//    │  left = Name (
		//    │  id = 'structure_type'
		//    │  ctx = Load (
		//    │
		//    │  ) // Load
		//    │  ) // Name
		//    │  ops =  [
		//    │  In (
		//    │
		//    │  ) // In
		//    │  ] //
		//    │  comparators =  [
		//    │  Attribute (
		//    │  value = Name (
		//    │  id = 'self'
		//    │  ctx = Load (
		//    │
		//    │  ) // Load
		//    │  ) // Name
		//    │  attr = 'horizontal_types'
		//    │  ctx = Load (
		//    │
		//    │  ) // Load
		//    │  ) // Attribute
		//    │  ] //
		//    │  ) // Compare
		//    │  Error in func transpileOp
		//    │  └──not valid token: In for In (
		//    │
		//    │  ) // In
		//    └──Nil binary: Compare (
		//       left = Name (
		//       id = 'structure_type'
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
		//       Attribute (
		//       value = Name (
		//       id = 'self'
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // Name
		//       attr = 'horizontal_types'
		//       ctx = Load (
		//
		//       ) // Load
		//       ) // Attribute
		//       ] //
		//       ) // Compare
		//
	}
	if self.load_condition == "slamming" {
		psl = self.__calculate_poly_value(0)
	} else {
		psl = 0
	}
	return max(press, psl)
}
func __calculate_poly_value() {
	"\n        Returning magnitude of load in the polynominal equation.\n        :param variable:\n        :return:\n        "
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
	//    │           Attribute (
	//    │           value = Name (
	//    │           id = 'self'
	//    │           ctx = Load (
	//    │
	//    │           ) // Load
	//    │           ) // Name
	//    │           attr = 'poly_third'
	//    │           ctx = Load (
	//    │
	//    │           ) // Load
	//    │           ) // Attribute
	//    │           Attribute (
	//    │           value = Name (
	//    │           id = 'self'
	//    │           ctx = Load (
	//    │
	//    │           ) // Load
	//    │           ) // Name
	//    │           attr = 'poly_second'
	//    │           ctx = Load (
	//    │
	//    │           ) // Load
	//    │           ) // Attribute
	//    │           Attribute (
	//    │           value = Name (
	//    │           id = 'self'
	//    │           ctx = Load (
	//    │
	//    │           ) // Load
	//    │           ) // Name
	//    │           attr = 'poly_first'
	//    │           ctx = Load (
	//    │
	//    │           ) // Load
	//    │           ) // Attribute
	//    │           Attribute (
	//    │           value = Name (
	//    │           id = 'self'
	//    │           ctx = Load (
	//    │
	//    │           ) // Load
	//    │           ) // Name
	//    │           attr = 'poly_const'
	//    │           ctx = Load (
	//    │
	//    │           ) // Load
	//    │           ) // Attribute
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
	//                Attribute (
	//                value = Name (
	//                id = 'self'
	//                ctx = Load (
	//
	//                ) // Load
	//                ) // Name
	//                attr = 'poly_third'
	//                ctx = Load (
	//
	//                ) // Load
	//                ) // Attribute
	//                Attribute (
	//                value = Name (
	//                id = 'self'
	//                ctx = Load (
	//
	//                ) // Load
	//                ) // Name
	//                attr = 'poly_second'
	//                ctx = Load (
	//
	//                ) // Load
	//                ) // Attribute
	//                Attribute (
	//                value = Name (
	//                id = 'self'
	//                ctx = Load (
	//
	//                ) // Load
	//                ) // Name
	//                attr = 'poly_first'
	//                ctx = Load (
	//
	//                ) // Load
	//                ) // Attribute
	//                Attribute (
	//                value = Name (
	//                id = 'self'
	//                ctx = Load (
	//
	//                ) // Load
	//                ) // Name
	//                attr = 'poly_const'
	//                ctx = Load (
	//
	//                ) // Load
	//                ) // Attribute
	//                ] //
	//                ctx = Load (
	//
	//                ) // Load
	//                ) // List
	//
}
func get_load_condition() {
	return self.load_condition
}
func is_tank_test() {
	return self.load_condition == "tanktest"
}
func get_load_parmeters() {
	return self.poly_third.self.poly_second.self.poly_first.self.poly_const.self.load_condition.None.self.manual_press.self.static_draft.self.name_of_load.self.limit_state
}
func get_name() {
	return self.name_of_load
}
func is_static() {
	"\n        Checking if the load is static type.\n        :return:\n        "
	return self.static_draft != None
}
func get_static_draft() {
	"\n        Return static draft if is_static\n        :return:\n        "
	if self.is_static() {
		return self.static_draft
	}
}
func get_limit_state() {
	" Return ULS, FLS.... "
	return self.limit_state
}
func get_load_condition() {
	" Getting loaded, ballast or part "
	return self.load_condition
}
func __init__() {
	self.properties = tank_dict
	self.compartment_number = tank_dict["comp_no"]
	self.cells = tank_dict["cells"]
	self.min_elevation = tank_dict["min_el"]
	self.max_elevation = tank_dict["max_el"]
	self.content = tank_dict["content"]
	self.added_pressure = tank_dict["added_press"]
	self.density = tank_dict["density"]
	self.acc_static = tank_dict["acc"]["static"]
	self.acc_dyn_loaded = tank_dict["acc"]["dyn_loaded"]
	self.acc_dyn_ballast = tank_dict["acc"]["dyn_ballast"]
	// Find PY4GO error
	// Error in func transpileStmt
	// └──Error in func transpileExprs
	//    ├──Error function: 5
	//    └──Name: List
	//       cannot transpile : List (
	//       elts =  [
	//       Str (
	//       s = 'crude_oil'
	//       ) // Str
	//       Str (
	//       s = 'diesel'
	//       ) // Str
	//       Str (
	//       s = 'slop'
	//       ) // Str
	//       Str (
	//       s = 'ballast'
	//       ) // Str
	//       ] //
	//       ctx = Load (
	//
	//       ) // Load
	//       ) // List
	//
}
func __str__() {
	"\n        Prints a string for the tank.\n        :return:\n        "
	tank_string = str("--- Tank properties (selected tank) ---" + "\n Minimum elevtaion:          " + str(self.min_elevation) + "\n Maximum elevation:          " + str(self.max_elevation) + "\n Content of tank:            " + self.content + "\n Defined density:            " + str(self.density) + "\n Defined acceleration:       " + "st = " + str(self.acc_static) + " , azl = " + str(self.acc_dyn_loaded) + " , azb = " + str(self.acc_dyn_ballast) + "\n Added pressure at tank top: " + str(self.added_pressure))
	return tank_string
}
func set_overpressure() {
	"\n        Setter\n        :param overpressure:\n        :return:\n        "
	self.added_pressure = overpressure
	self.properties["added_press"] = overpressure
}
func set_content() {
	"\n        Setter\n        :param overpressure:\n        :return:\n        "
	self.properties["content"] = content
	self.content = content
}
func set_acceleration() {
	"\n        Setter\n        :param overpressure:\n        :return:\n        "
	self.acc_static = acc["static"]
	self.properties["static"] = acc["static"]
	self.acc_dyn_loaded = acc["dyn_loaded"]
	self.properties["dyn_loaded"] = acc["dyn_loaded"]
	self.acc_dyn_ballast = acc["dyn_ballast"]
	self.properties["dyn_ballast"] = acc["dyn_ballast"]
}
func set_density() {
	"\n        Setter\n        :param overpressure:\n        :return:\n        "
	self.properties["density"] = density
	self.density = density
}
func get_name() {
	"\n        Returns the name of the compartmnet\n        :return:\n        "
	return "comp" + str(self.compartment_number)
}
func get_highest_elevation() {
	"\n        Find the top of the tank.\n        :return:\n        "
	return self.max_elevation
}
func get_lowest_elevation() {
	"\n        Find the bottom of the tank.\n        :return:\n        "
	return self.min_elevation
}
func get_line_pressure_from_max_pressure() {
	"\n        Used when you have a maximum pressure and request the pressure at a specific coordinate.\n        :param coordinates:\n        :return:\n        "
	elevation = coordinates[1]
	return pressure * ((self.get_highest_elevation() - elevation) / (self.get_highest_elevation() - self.get_lowest_elevation()))
}
func get_calculated_pressure() {
	"\n        Get the pressure with specified variable.\n        :param elevaiton:\n        :return:\n        "
	elevation = coordinates[1]
	press = (self.get_highest_elevation() - elevation) * self.density * acceleration
	return press
}
func get_bottom_pressure() {
	"\n        Get pressure at bottom of tank.\n        :return:\n        "
	return (self.get_highest_elevation()-self.get_lowest_elevation())*self.density*self.acceleration + self.added_pressure
}
func get_top_pressure() {
	"\n        Get the pressure at the top of the tank.\n        :return:\n        "
	return self.added_pressure
}
func get_density() {
	"\n        Get the tank density.\n        :return:\n        "
	return self.density
}
func get_content() {
	"\n        Returnt the tank content type\n        :return:\n        "
	return self.content
}
func get_accelerations() {
	"\n        Returns the defined accelerations\n        :return:\n        "
	return self.acc_static.self.acc_dyn_loaded.self.acc_dyn_ballast
}
func get_overpressure() {
	"\n        Get the overpressure at tank top.\n        :return:\n        "
	return self.added_pressure
}
func get_parameters() {
	"\n        Returns properties\n        :return:\n        "
	return self.properties
}
func is_loaded_condition() {
	""\n        Check to see if the tank shall be in cluded in loaded condition.\n        self.tank_options = [crude_oil, diesel, slop, ballast]\n        :return:\n        ""
}
func is_ballast_condition() {
	"\n        Check to see if the tank shall be in cluded in loaded condition.\n        :return:\n        "
}
func is_tank_test_condition() {
	"\n        Check to see if the tank shall be in cluded in loaded condition.\n        :return:\n        "
}
func get_condition() {
	""\n        Returning the condition.\n        self.load_conditions = [loaded, ballast,tanktest]\n        :return:\n        ""
}
func get_tank_dnv_minimum_pressure() {
	"\n        Calculating 4.3.7 and 4.3.8 and returning the highest of these pressures.\n        :return:\n        "
	if self.is_loaded_condition() {
		dyn_acc = self.acc_dyn_loaded
	} else {
		if self.is_ballast_condition() {
			dyn_acc = self.acc_dyn_ballast
		} else {
			dyn_acc = 0
		}
	}
	hop = self.get_highest_elevation() - self.get_lowest_elevation()
	p_4_3_7 = self.density * self.acc_static * hop * (lf_static + dyn_acc/self.acc_static*lf_enviromental)
	p_4_3_8 = (self.density*self.acc_static*hop + self.get_overpressure()) * lf_static
	return max(p_4_3_7, p_4_3_8)
}
func __init__() {
	"\n        Input from main application is:\n        line for this object\n        tank_dict = {} #main tank dictionary (created when BFS search is executed for the grid) (comp# : TankObj)\n        load_dict = {} #main load dictionary (created in separate load window (load# : [LoadObj, lines])\n        comb_dict = {} #load combination dictionary (comb,line,load) : [DoubleVar(), DoubleVar], IntVar()]\n        "
	self.object_line = object_line
	self.comb_dict = comb_dict
	self.tank_dict = tank_dict
	self.load_dict = load_dict
}
func __str__() {
	return "NOT IMPLEMENTED"
}
func get_load_factors() {
	"\n        Get the tk.DoubleVar, tk.DoubleVar, tk.IntVar that is used in the load factor input and on/off.\n        :return:\n        "
	return self.load_factor_static.get()[self.load_factor_dynamic.get()][self.on_off.get()]
}
func get_load_factor_static() {
	"\n        Setting the the dynamic load factor.\n        :return:\n        "
	return self.load_factor_static.get()
}
func get_load_factor_dynamic() {
	"\n        Setting the the dynamic load factor.\n        :return:\n        "
	return self.load_factor_dynamic.get()
}
func get_on_off() {
	"\n        Setting the the dynamic load factor.\n        :return:\n        "
	return self.on_off.get()
}
func set_load_factor_static() {
	"\n        Setting the the dynamic load factor.\n        :return:\n        "
	self.load_factor_static = value
}
func set_load_factor_dynamic() {
	"\n        Setting the the dynamic load factor.\n        :return:\n        "
	self.load_factor_dynamic = value
}
func set_on_off() {
	"\n        Setting the the dynamic load factor.\n        :return:\n        "
	self.on_off = value
}
func set_combination_dictionary() {
	"\n        Setting the combination dictionary.\n        :return:\n        "
	self.comb_dict = value
}
func set_load_dictionary() {
	"\n        Setting the load dictionary.\n        :return:\n        "
	self.load_dict = value
}
func set_tank_dictionary() {
	"\n        Setting the tank dictionary.\n        :return:\n        "
	self.tank_dict = value
}
func main() {
	"\n    This Class calculates the load to be applied on the structure\n    "
	""\n    This class incorporates all tank definitions\n        temp_tank_dict = {0  comp_no : comp_no,\n                          1  cells : properties[0],\n                          2  min_el :  properties[1],\n                          3  max_el : properties[2],\n                          4  content : ,\n                          5  added_press : 0,\n                          6  acc : {static:g,dyn_loaded:az,dyn_ballast:az}\n                          7  density : 1025}\n\n    ""
	"\n    THIS CLASS IS CURRENTLY NOT USED. MAY NOT BE USED AT ALL. IT IS STUPID.\n    This class cointaines the load combinations.   \n    combination,self.active_line,compartment\n    "
	if __name__ == "__main__" {
		// Find PY4GO error
		// Error in func transpileStmt
		// └──Error in func transpileExprs
		//    └──Error in func transpileExprs
		//       ├──Error in func transpileExprs
		//       │  ├──Error function: 5
		//       │  └──Name: List
		//       │     cannot transpile : List (
		//       │     elts =  [
		//       │     Call (
		//       │     func = Name (
		//       │     id = 'Loads'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Name
		//       │     args =  [
		//       │     Attribute (
		//       │     value = Name (
		//       │     id = 'ex'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Name
		//       │     attr = 'load_bottom'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Attribute
		//       │     ] //
		//       │     keywords =  [
		//       │
		//       │     ] //
		//       │     starargs = None
		//       │     kwargs = None
		//       │     ) // Call
		//       │     Call (
		//       │     func = Name (
		//       │     id = 'Loads'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Name
		//       │     args =  [
		//       │     Attribute (
		//       │     value = Name (
		//       │     id = 'ex'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Name
		//       │     attr = 'load_side'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Attribute
		//       │     ] //
		//       │     keywords =  [
		//       │
		//       │     ] //
		//       │     starargs = None
		//       │     kwargs = None
		//       │     ) // Call
		//       │     Call (
		//       │     func = Name (
		//       │     id = 'Loads'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Name
		//       │     args =  [
		//       │     Attribute (
		//       │     value = Name (
		//       │     id = 'ex'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Name
		//       │     attr = 'load_static'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Attribute
		//       │     ] //
		//       │     keywords =  [
		//       │
		//       │     ] //
		//       │     starargs = None
		//       │     kwargs = None
		//       │     ) // Call
		//       │     Call (
		//       │     func = Name (
		//       │     id = 'Loads'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Name
		//       │     args =  [
		//       │     Attribute (
		//       │     value = Name (
		//       │     id = 'ex'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Name
		//       │     attr = 'load_slamming'
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // Attribute
		//       │     ] //
		//       │     keywords =  [
		//       │
		//       │     ] //
		//       │     starargs = None
		//       │     kwargs = None
		//       │     ) // Call
		//       │     ] //
		//       │     ctx = Load (
		//       │
		//       │     ) // Load
		//       │     ) // List
		//       └──Error in func transpileExprs
		//          ├──Error function: 5
		//          └──Name: List
		//             cannot transpile : List (
		//             elts =  [
		//             Str (
		//             s = 'BOTTOM'
		//             ) // Str
		//             Str (
		//             s = 'SIDE_SHELL'
		//             ) // Str
		//             Str (
		//             s = ''
		//             ) // Str
		//             Str (
		//             s = ''
		//             ) // Str
		//             ] //
		//             ctx = Load (
		//
		//             ) // Load
		//             ) // List
		//
	}
}
