package org.azzgo.designPattern.factoryMethod;

import org.azzgo.designPattern.factoryMethod.framework.*;
import org.azzgo.designPattern.factoryMethod.idcard.IDCardFactory;
import org.junit.Test;

public class factoryMethodTest {
    @Test()
    public void TestFactoryMethod() {
        ProductFactory factory = new IDCardFactory();
        Product card1 = factory.create("小明");
        Product card2 = factory.create("红红");
        Product card3 = factory.create("蕾蕾");

        card1.use();
        card2.use();
        card3.use();
    }
}

