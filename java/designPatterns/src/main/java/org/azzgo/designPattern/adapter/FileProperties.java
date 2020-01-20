package org.azzgo.designPattern.adapter;

import java.io.*;
import java.util.Properties;

public class FileProperties implements FileIO {
    Properties properties = new Properties();;

    @Override
    public void readFromFile(String filename) throws IOException {
        InputStream fileStream = new FileInputStream(filename);
        properties.load(fileStream);
    }

    @Override
    public void writeToFile(String filename) throws IOException {
        properties.store(new FileWriter(filename), "written by FileProperties.");
    }

    @Override
    public void setValue(String key, String value) {
        properties.setProperty(key, value);
    }

    @Override
    public String getValue(String key) {
        return properties.getProperty(key);
    }
}
