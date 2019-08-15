package org.azzgo.designPattern.state;

import org.azzgo.designPattern.state.states.Signal;

public class JorgeAvenueTrafficLight implements TrafficLight {

    private Signal signal;

    @Override
    public void nextSignal() {
        if (signal == null) {
            System.out.println("红绿灯损坏中~~~~~~~~");
        }

        System.out.println(signal + "，红绿灯切换中。。。");
        signal.nextSignal(this);
        System.out.println(signal);
        System.out.println(signal.passable());
    }

    @Override
    public void setSignal(Signal signal) {
        this.signal = signal;
    }
}
