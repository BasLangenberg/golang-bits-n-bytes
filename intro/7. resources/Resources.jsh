System.out.println(" ===== ")

FileInputStream fis1 = null;
try {
    fis1 = new FileInputStream(new File("./Resources.jsh"));

    int i; 
            
    while ((i = fis1.read()) != -1) {
        System.out.print((char) i);
    }
} finally {
    if (fis1 != null) {
        fis1.close();
    }
}

System.out.println(" ===== ")

try (FileInputStream fis2 = new FileInputStream(new File("./Resources.jsh"))) {
    int i; 
            
    while ((i = fis2.read()) != -1) {
        System.out.print((char) i);
    }
}

System.out.println(" ===== ")