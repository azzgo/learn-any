package org.azzgo.designPattern.factoryMethod.framework;

public abstract class ProductFactory {
    public final Product create(String owner) {
        Product p =  createProduct(owner);
        registerProduct(p);
        return p;
    };

    protected abstract Product createProduct(String owner);
    protected abstract void registerProduct(Product product);

}
