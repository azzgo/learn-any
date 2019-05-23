package org.azzgo.designPattern.singleton;

import org.junit.Assert;
import org.junit.Test;

public class SimpleSingletonTest {

    @Test public void singletonTestShouldWork() {
        SimpleSingleton simpleSingleton1 = SimpleSingleton.getInstance();
        SimpleSingleton simpleSingleton2 = SimpleSingleton.getInstance();

        Assert.assertEquals(null, simpleSingleton1, simpleSingleton2);
    }
}
