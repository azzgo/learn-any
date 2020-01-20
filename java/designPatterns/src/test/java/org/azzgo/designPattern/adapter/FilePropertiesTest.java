package org.azzgo.designPattern.adapter;

import org.junit.Assert;
import org.junit.Test;

import java.io.IOException;

public class FilePropertiesTest {
    @Test public void playPropensities() {
        try {
            FileIO f = new FileProperties();
            f.readFromFile(getClass().getClassLoader().getResource("file.txt").getPath());
            f.setValue("year", "2004");
            f.setValue("month", "4");
            f.setValue("day", "21");
            f.writeToFile("/var/tmp/newFile.txt");


            FileIO newF = new FileProperties();

            newF.readFromFile("/var/tmp/newFile.txt");

            Assert.assertEquals(newF.getValue("year"), "2004");
            Assert.assertEquals(newF.getValue("month"), "4");
            Assert.assertEquals(newF.getValue("day"), "21");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
