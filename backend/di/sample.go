package di

import "backend/presentation/rest/controller"

func Sample() controller.Sample {
	return controller.NewSample()
}
