package templates

func Repository() string {
	return `
import { PrismaClient } from "@prisma/client";
import ResourceRepository from "sygnus";
import prismaClient from "../database/prismaClient";

export default class {{.PCaseName}}Repository extends BaseRepository {
  constructor() {
    super(prismaClient, "{{.Name}}"); 
  } 
}

  `
}
