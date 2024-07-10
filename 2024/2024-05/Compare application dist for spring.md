# Compare application dist for spring

```java
import java.io.File;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class JarComparator {

    public static void main(String[] args) {
        String[] build1Dirs = {"build/install/service/lib", "build/libs"};
        String[] build2Dirs = {"build2/install/service/lib", "build2/libs"};

        List<String> missingJars = compareDirectories(build1Dirs, build2Dirs);

        if (missingJars.isEmpty()) {
            System.out.println("No missing jars found.");
        } else {
            System.out.println("Missing jars:");
            for (String jar : missingJars) {
                System.out.println(jar);
            }
        }
    }

    private static List<String> compareDirectories(String[] build1Dirs, String[] build2Dirs) {
        List<File> jars1 = getJarFiles(build1Dirs);
        List<File> jars2 = getJarFiles(build2Dirs);

        List<String> missingJars = new ArrayList<>();

        for (File jar : jars1) {
            if (!containsFile(jars2, jar)) {
                missingJars.add(jar.getPath());
            }
        }

        for (File jar : jars2) {
            if (!containsFile(jars1, jar)) {
                missingJars.add(jar.getPath());
            }
        }

        return missingJars;
    }

    private static List<File> getJarFiles(String[] dirs) {
        List<File> jarFiles = new ArrayList<>();

        for (String dir : dirs) {
            File directory = new File(dir);
            if (directory.exists() && directory.isDirectory()) {
                File[] files = directory.listFiles();
                if (files != null) {
                    jarFiles.addAll(Arrays.asList(files));
                }
            }
        }

        return jarFiles;
    }

    private static boolean containsFile(List<File> files, File target) {
        for (File file : files) {
            if (file.getName().equals(target.getName())) {
                return true;
            }
        }
        return false;
    }
}
```

#jar #build #spring
#draft