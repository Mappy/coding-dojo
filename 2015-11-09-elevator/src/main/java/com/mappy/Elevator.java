package com.mappy;

public interface Elevator {
    void reset();

    void userHasEntered();

    void userHasExited();

    void call(String atFloor, String to);

    void go(String floorToGo);

    String nextCommand();
}