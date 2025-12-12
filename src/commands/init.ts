 import { mkdir, writeFile, access } from "node:fs/promises";
 import { join } from "node:path";

 export async function initRepository() {
   const root = process.cwd();
   const karasuPath = join(root, ".karasu");

   try {
     await access(karasuPath);
     console.log("Repository already exists");
     return;
   } catch {
     // Repository does not exist, proceed
   }

   // Create structure
   await mkdir(join(karasuPath, "objects"), { recursive: true });
   await mkdir(join(karasuPath, "refs", "heads"), { recursive: true });

   // HEAD -> current branch
   await writeFile(
     join(karasuPath, "HEAD"),
     "ref: refs/heads/main\n",
     "utf-8"
   );

   // Create main branch empty
   await writeFile(
     join(karasuPath, "refs", "heads", "main"),
     "",
     "utf-8"
   );

   // Create index empty
   await writeFile(join(karasuPath, "index"), "", "utf-8");

   console.log("Karasu Repository Initialized");
   console.log(`DIR: ${karasuPath}`);
 }
