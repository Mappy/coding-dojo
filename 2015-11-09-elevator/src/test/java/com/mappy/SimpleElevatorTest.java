package com.mappy;

import org.junit.Test;

import static org.assertj.core.api.Assertions.*;

public class SimpleElevatorTest {
    private final Elevator elevator = new SimpleElevator();

    @Test
    public void should_be_true() throws Exception {
        assertThat(elevator.nextCommand()).isEqualTo("OPEN");
        assertThat(elevator.nextCommand()).isEqualTo("CLOSE");
        assertThat(elevator.nextCommand()).isEqualTo("UP");
    }

    @Test
    public void should_be_true_with_stage() throws Exception {
        for (int i = 0; i < 5; i++) {
            assertThat(elevator.nextCommand()).isEqualTo("OPEN");
            assertThat(elevator.nextCommand()).isEqualTo("CLOSE");
            assertThat(elevator.nextCommand()).isEqualTo("UP");
        }
    }

    @Test
    public void should_return_to_0_after_stage_5() {
        for (int i = 0; i < 5; i++) {
            assertThat(elevator.nextCommand()).isEqualTo("OPEN");
            assertThat(elevator.nextCommand()).isEqualTo("CLOSE");
            assertThat(elevator.nextCommand()).isEqualTo("UP");
        }
        assertThat(elevator.nextCommand()).isEqualTo("OPEN");
        assertThat(elevator.nextCommand()).isEqualTo("CLOSE");
        assertThat(elevator.nextCommand()).isEqualTo("DOWN");
        assertThat(elevator.nextCommand()).isEqualTo("DOWN");
        assertThat(elevator.nextCommand()).isEqualTo("DOWN");
        assertThat(elevator.nextCommand()).isEqualTo("DOWN");
        assertThat(elevator.nextCommand()).isEqualTo("DOWN");
        assertThat(elevator.nextCommand()).isEqualTo("OPEN");
    }
}
