package org.azzgo.designPattern.proxy;

public interface Printable {
    void setPrinterName(String name);

    String getPrinterName();

    void print(String string);
}
