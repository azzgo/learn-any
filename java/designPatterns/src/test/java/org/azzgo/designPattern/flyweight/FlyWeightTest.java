package org.azzgo.designPattern.flyweight;

import org.junit.Test;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

public class FlyWeightTest {
    @Test
    public void work() {
        String testString = "15688887777";

        List<BigChar> bigChars = new ArrayList<>();
        BigCharFactory bigCharFactory = BigCharFactory.getFactoryInstance();

        for (char c : testString.toCharArray()) {
            bigChars.add(bigCharFactory.getBigChar(c));
        }

        for (BigChar bigChar : bigChars) {
            bigChar.print();
        }
    }
}
