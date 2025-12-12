#!/usr/bin/env bun

import { Command } from "commander";
import { initRepository } from "./commands/init";

const program = new Command();

program
  .name("karasu")
  .description("Karasu VCS - un sistema de control de versiones minimalista")
  .version("0.1.0");

program
  .command("init")
  .description("Inicializa un nuevo repositorio Karasu en este directorio")
  .action(async () => {
    await initRepository();
  });

program.parse();

