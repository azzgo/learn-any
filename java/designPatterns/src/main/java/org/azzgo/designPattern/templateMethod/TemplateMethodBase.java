package org.azzgo.designPattern.templateMethod;

public abstract class TemplateMethodBase {
    abstract String do1();
    abstract String do2();

    String templateMethod() {
        return "TemplateMethod" + this.do1() + this.do2();
    }
}
