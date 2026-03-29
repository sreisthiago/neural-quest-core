use wasm_bindgen::prelude::*;
use web_sys::Performance;

#[wasm_bindgen]
pub struct NeuralEngine {
    perf: Performance,
    last_start: f64,
}

#[wasm_bindgen]
impl NeuralEngine {
    pub fn new() -> Result<NeuralEngine, JsValue> {
        let window = web_sys::window().ok_or("No window")?;
        let perf = window.performance().ok_or("No performance API")?;
        Ok(NeuralEngine { perf, last_start: 0.0 })
    }

    pub fn trigger_stimulus(&mut self) {
        self.last_start = self.perf.now();
    }

    pub fn get_latency(&self) -> f64 {
        self.perf.now() - self.last_start
    }
}
