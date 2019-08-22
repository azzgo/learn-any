package org.azzgo.designPattern.flyweight;

import java.util.HashMap;
import java.util.Map;

public class BigCharFactory {
    private Map<Character, BigChar> pool = new HashMap<>();
    static private BigCharFactory factory = new BigCharFactory();

    private BigCharFactory() {}

    static BigCharFactory getFactoryInstance() {
        return factory;
    }

    public synchronized BigChar getBigChar(char charName) {
        BigChar bigChar = pool.get(charName);

        if (bigChar == null) {
            bigChar = new BigChar(charName);
            pool.put(charName, bigChar);
        }

        return bigChar;
    }
}
