package templates

func Repository() string {
	return `
import { PrismaClient } from "@prisma/client";
import { ResourceRepository } from "sygnus";
import prismaClient from "../database/prismaClient";

export default class {{.PCaseName}}Repository extends ResourceRepository {
  constructor() {
    super(prismaClient, "{{.Name}}"); 
  } 
}

  `
}
