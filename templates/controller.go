package templates

func Controller() string {
	return `
import {{.PCaseName}}UseCase from "../useCases/{{.PCaseName}}UseCase";
import { ResourceController } from "sygnus";

export default class {{.PCaseName}}Controller extends ResourceController {
  constructor() {
    super(new {{.PCaseName}}UseCase());
  }
}
  
  `
}
