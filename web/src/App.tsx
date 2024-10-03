import { createSignal } from "solid-js";
import "./App.css";

function App() {
	const [count, setCount] = createSignal(0);

	return (
		<>
			<div>
				<h1>Hello Jager!</h1>
			</div>
			<h1>Vite + Solid</h1>
			<div class="card">
				<button onClick={() => setCount((count) => count + 1)}>
					count is {count()}
				</button>
				<p>
					Edit <code>src/App.tsx</code> and save to test HMR
				</p>
			</div>
			<p class="read-the-docs">
				Click on the Vite and Solid logos to learn more
			</p>
		</>
	);
}

export default App;
