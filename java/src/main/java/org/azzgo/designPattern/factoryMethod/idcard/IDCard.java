package org.azzgo.designPattern.factoryMethod.idcard;

import org.azzgo.designPattern.factoryMethod.framework.Product;

public class IDCard extends Product {
    private final String owner;
    private final String cardId;

    IDCard(String owner, String cardId) {
        System.out.println("制作" + owner + "的 ID 卡, 制作卡号是" + cardId);
        this.owner = owner;
        this.cardId = cardId;
    }

    @Override()
    public void use() {
        System.out.println("使用" + owner + "的 ID 卡, " + "卡号是" + cardId);
    }

    public String getOwner() {
        return this.owner;
    }

    public String getCardId() {
        return cardId;
    }
}
