package templates

func UseCase() string {
	return `
import {{.PCaseName}}Repository from "../repositories/{{.PCaseName}}Repository";
import { ResourceUseCase } from "sygnus";

export default class {{.PCaseName}}UseCase extends ResourceUseCase<{{.PCaseName}}Repository> {
  constructor() {
    super(new {{.PCaseName}}Repository());
  }
}

  `
}
