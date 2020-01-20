package org.azzgo.designPattern.strategy;

public class AirPlane implements Vehicle {

    @Override
    public void transition() {
        System.out.println("transition on airplane！");
    }
}
