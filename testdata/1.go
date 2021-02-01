package main

func __init__() {
	self.tol = 0.0001
	self.optiter = 5
	self.iterMax = 10
	self.fixedStep = False
	self.maxFactor = 1e+20
	globdat.totalFactor = 1.0
	globdat.factor = 1.0
	self.maxLam = 1e+20
	dofCount = len(globdat.dofs)
	BaseModule.__init__(self, props)
	if !hasattr(globdat, "Daprev") {
		globdat.Daprev = zeros(dofCount)
		globdat.Dlamprev = 1.0
	}
	globdat.lam = 1.0
	Print("\n  Starting Riks arclength solver\n")
}
func run() {
	stat = globdat.solverStatus
	stat.increaseStep()
	a = globdat.state
	Da = globdat.Dstate
	fhat = globdat.fhat
	self.printHeader(stat.cycle)
	error = 1.0
	if stat.cycle == 1 {
		K, fint = assembleTangentStiffness(props, globdat)
		Da1 = globdat.dofs.solve(K, globdat.lam*fhat)
		Dlam1 = globdat.lam
	} else {
		Da1 = globdat.factor * globdat.Daprev
		Dlam1 = globdat.factor * globdat.Dlamprev
		globdat.lam = globdat.lam + Dlam1
	}
	a[:] = a[:] + Da1[:]
	Da[:] = Da1[:]
	Dlam = Dlam1
	K, fint = assembleTangentStiffness(props, globdat)
	res = globdat.lam*fhat - fint
	for error > self.tol {
		stat.iiter = stat.iiter + 1
		d1 = globdat.dofs.solve(K, fhat)
		d2 = globdat.dofs.solve(K, res)
		ddlam = -dot(Da1, d2) / dot(Da1, d1)
		dda = ddlam*d1 + d2
		Dlam = Dlam + ddlam
		globdat.lam = globdat.lam + ddlam
		Da[:] = Da[:] + dda[:]
		a[:] = a[:] + dda[:]
		K, fint = assembleTangentStiffness(props, globdat)
		res = globdat.lam*fhat - fint
		error = globdat.dofs.norm(res) / globdat.dofs.norm(globdat.lam*fhat)
		self.printIteration(stat.iiter, error)
		if stat.iiter == self.iterMax {
			RuntimeError("Newton-Raphson iterations did not converge!")
		}
	}
	self.printConverged(stat.iiter)
	globdat.elements.commitHistory()
	globdat.fint = fint
	if !self.fixedStep {
		globdat.factor = pow(0.5, 0.25*(stat.iiter-self.optiter))
		globdat.totalFactor = globdat.totalFactor * globdat.factor
	}
	if globdat.totalFactor > self.maxFactor {
		globdat.factor = 1.0
	}
	globdat.Daprev[:] = Da[:]
	globdat.Dlamprev = Dlam
	if globdat.lam > self.maxLam || stat.cycle > 1000 || a[globdat.dofs.getForType(4, "v")] > 5 {
		globdat.active = False
	}
}
func printHeader() {
	Print("\n======================================")
	Print(" Load step %i" + cycle)
	Print("======================================")
	Print("  iter # : L2-norm residual")
}
func printIteration() {
	Print("   %5i : %4.2e " + iiter.error)
}
func printConverged() {
	Print("--------------------------------------")
	Print(" Converged in %i iterations" + iiter)
}
