package org.azzgo.designPattern.factoryMethod.idcard;

import org.azzgo.designPattern.factoryMethod.framework.Product;
import org.azzgo.designPattern.factoryMethod.framework.ProductFactory;

import java.util.*;

public class IDCardFactory extends ProductFactory {
    private Map owners = new HashMap<String, String>();

    @Override
    protected Product createProduct(String owner) {
        UUID uuid = UUID.randomUUID();
        return new IDCard(owner, uuid.toString());
    }

    @Override
    protected void registerProduct(Product product) {
        IDCard idCard = (IDCard) product;
        this.owners.put(idCard.getCardId(), idCard.getOwner());
    }
}
