package main

var Polynomials map[string]Polynomial = map[string]Polynomial{
	"270-600/12V": Polynomial{
		Gain:                    1.0 / 100.0,
		ApproximationType:       "MACLAURIN",
		FrequencyLowerBound:     0.0,
		FrequencyUpperBound:     0.0,
		ApproximationLowerBound: 600.0,
		ApproximationUpperBound: 1100.0,
		MaximumError:            0.0,
		Coefficients:            []Coefficient{Coefficient{600.0}, Coefficient{100.0}},
	},
	"270-600/24V": Polynomial{
		Gain:                    1.0 / 100.0,
		ApproximationType:       "MACLAURIN",
		FrequencyLowerBound:     0.0,
		FrequencyUpperBound:     0.0,
		ApproximationLowerBound: 600.0,
		ApproximationUpperBound: 1100.0,
		MaximumError:            0.0,
		Coefficients:            []Coefficient{Coefficient{600.0}, Coefficient{100.0}},
	},
	"270-800/12V": Polynomial{
		Gain:                    1.0 / 60.0,
		ApproximationType:       "MACLAURIN",
		FrequencyLowerBound:     0.0,
		FrequencyUpperBound:     0.0,
		ApproximationLowerBound: 800.0,
		ApproximationUpperBound: 1100.0,
		MaximumError:            0.0,
		Coefficients:            []Coefficient{Coefficient{800.0}, Coefficient{100.0}},
	},
	"Druck PTX-1830-A": Polynomial{
		Gain:                    0.8 / 1000.0,
		ApproximationType:       "MACLAURIN",
		FrequencyLowerBound:     0.0,
		FrequencyUpperBound:     0.0,
		ApproximationLowerBound: 0.0,
		ApproximationUpperBound: 20.0,
		MaximumError:            0.0,
		Coefficients:            []Coefficient{Coefficient{4.0 / 1000.0}, Coefficient{0.8 / 1000.0}},
	},
	"Druck PTX-1830-B": Polynomial{
		Gain:                    250.0,
		ApproximationType:       "MACLAURIN",
		FrequencyLowerBound:     0.0,
		FrequencyUpperBound:     0.0,
		ApproximationLowerBound: 0.004,
		ApproximationUpperBound: 0.02,
		MaximumError:            0.0,
		Coefficients:            []Coefficient{Coefficient{0.0}, Coefficient{250.0}},
	},
	"InfraBSU microphone": Polynomial{
		Gain:                    0.008,
		ApproximationType:       "MACLAURIN",
		FrequencyLowerBound:     0.0,
		FrequencyUpperBound:     0.0,
		ApproximationLowerBound: 0.0,
		ApproximationUpperBound: 0.0,
		MaximumError:            0.0,
		Coefficients:            []Coefficient{Coefficient{0.0}, Coefficient{125.0}},
		Notes:                   &[]string{"this is incorrect, enough to get a plot\nin theory this is linear with +/- 10mV for +/- 125Pa (1.25 mbar)\n which translates to +/- 1V for +/- 125mbar\n"}[0],
	},
	"LM35": Polynomial{
		Gain:                    10.0 / 1000.0,
		ApproximationType:       "MACLAURIN",
		FrequencyLowerBound:     0.0,
		FrequencyUpperBound:     0.0,
		ApproximationLowerBound: -55.0,
		ApproximationUpperBound: 150.0,
		MaximumError:            0.0,
		Coefficients:            []Coefficient{Coefficient{0.0}, Coefficient{100.0}},
	},
	"SEAWATER": Polynomial{
		Gain:                    1030.0 / 1000.0,
		ApproximationType:       "MACLAURIN",
		FrequencyLowerBound:     0.0,
		FrequencyUpperBound:     0.0,
		ApproximationLowerBound: 0.0,
		ApproximationUpperBound: 20.0,
		MaximumError:            0.0,
		Coefficients:            []Coefficient{Coefficient{0.0}, Coefficient{1030.0 / 1000.0}},
	},
}
