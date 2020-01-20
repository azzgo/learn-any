package org.azzgo.designPattern.flyweight;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class BigChar implements IBigChar {
    private String fontData;

    BigChar(char charName) {
        try {
            BufferedReader reader = new BufferedReader(new FileReader(
                    getClass().getClassLoader().getResource("bigChars/big" + charName + ".txt").getPath()
            ));
            String line;
            StringBuilder buf = new StringBuilder();

            while ((line = reader.readLine()) != null) {
                buf.append(line);
                buf.append("\n");
            }
            reader.close();
            this.fontData = buf.toString();
        } catch (IOException e) {
            this.fontData = charName + "?";
        }
    }


    @Override
    public void print() {
        System.out.print(fontData);
    }
}
