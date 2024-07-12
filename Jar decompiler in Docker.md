Декомпилирование `.class` файлов внутри `jar` файлов для сравнения этих самых `jar` файлов. Использует [cfr](https://www.benf.org/other/cfr/) для декомпилирования. 

## Build

```bash
docker build -t decompiler .
```

## Usage

```bash
docker run --rm -v /path/to/jars:/jars decompiler /jars/SomeJar.jar
```

## Dockerfile

```dockerfile
####
# Build image
####
FROM openjdk:21 AS build
LABEL maintainer=avvero

RUN microdnf install findutils

COPY gradlew /app/
COPY gradle /app/gradle
WORKDIR /app
RUN ./gradlew --version

WORKDIR /app
COPY . .
RUN ./gradlew build --no-daemon

####
# Runtime image
####
FROM openjdk:21

COPY --from=build /app/build/classes/java/main/Decompiler.class /app/Decompiler.class
COPY --from=build /app/cfr.jar /app/cfr.jar

WORKDIR /app

ENTRYPOINT ["java", "Decompiler"]
```

## Code
```java
import java.io.*;
import java.nio.file.*;
import java.util.*;
import java.util.jar.JarEntry;
import java.util.jar.JarFile;

public class Decompiler {

    private static final Set<String> IGNORED_FILES = new HashSet<>(Arrays.asList(
            "META-INF/MANIFEST.MF"
    ));

    public static void main(String[] args) throws IOException {
        if (args.length == 0) {
            System.out.println("Usage: java Decompiler <path-to-jar1> <path-to-jar2> ...");
            System.exit(1);
        }

        for (String jarFilePath : args) {
            File jarFile = new File(jarFilePath);
            if (!jarFile.exists() || !jarFile.isFile()) {
                System.err.println("File not found or not a file: " + jarFilePath);
                continue;
            }

            String baseName = jarFile.getName().substring(0, jarFile.getName().lastIndexOf('.'));
            Path outputDir = jarFile.getParentFile().toPath().resolve(baseName + "_decompiled");
            Files.createDirectories(outputDir);

            try (JarFile jar = new JarFile(jarFile)) {
                decompileJar(jar, outputDir);
            } catch (IOException e) {
                System.err.println("Failed to decompile JAR: " + jarFilePath);
                e.printStackTrace();
            }
        }
    }

    private static void decompileJar(JarFile jarFile, Path outputDir) throws IOException {
        Enumeration<JarEntry> entries = jarFile.entries();
        while (entries.hasMoreElements()) {
            JarEntry entry = entries.nextElement();
            if (IGNORED_FILES.contains(entry.getName())) {
                continue;
            }

            File entryFile = new File(outputDir.toFile(), entry.getName());
            if (entry.isDirectory()) {
                entryFile.mkdirs();
            } else if (entry.getName().endsWith(".class")) {
                entryFile.getParentFile().mkdirs();
                Files.copy(jarFile.getInputStream(entry), entryFile.toPath(), StandardCopyOption.REPLACE_EXISTING);
                decompileClassFile(entryFile.toPath());
                Files.delete(entryFile.toPath());
            } else {
                entryFile.getParentFile().mkdirs();
                Files.copy(jarFile.getInputStream(entry), entryFile.toPath(), StandardCopyOption.REPLACE_EXISTING);
            }
        }
    }

    private static void decompileClassFile(Path classFile) throws IOException {
        Path outputDir = classFile.getParent();
        ProcessBuilder processBuilder = new ProcessBuilder("java", "-jar", "cfr.jar", classFile.toString(), "--outputdir", outputDir.toString());
        processBuilder.redirectErrorStream(true);
        Process process = processBuilder.start();

        try (BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()))) {
            String line;
            while ((line = reader.readLine()) != null) {
                System.out.println(line);
            }
            process.waitFor();
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            throw new IOException("Decompilation interrupted", e);
        }
    }
}
```

#java #jar #decompile #docker