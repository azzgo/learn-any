package org.azzgo.designPattern.state.states;

import org.azzgo.designPattern.state.TrafficLight;

public class RedSignal implements Signal {
    @Override
    public void nextSignal(TrafficLight context) {
        context.setSignal(new GreenSignal());
    }

    @Override
    public String toString() {
        return "红灯";
    }

    @Override
    public String passable() {
        return "当前红灯，你不能通过这里！";
    }
}
