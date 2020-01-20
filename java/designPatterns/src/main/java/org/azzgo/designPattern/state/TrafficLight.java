package org.azzgo.designPattern.state;

import org.azzgo.designPattern.state.states.Signal;

public interface TrafficLight {
    void nextSignal();

    void setSignal(Signal signal);
}
