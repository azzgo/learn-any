package org.azzgo.designPattern.templateMethod;

import org.junit.Assert;
import org.junit.Test;

public class TemplateMethodTest {
    @Test public void AssertMethodSame() {
        Implement1 implement1 = new Implement1();
        Implement2 implement2 = new Implement2();

        Assert.assertNotEquals(implement1.templateMethod(), implement2.templateMethod());
    }
}
