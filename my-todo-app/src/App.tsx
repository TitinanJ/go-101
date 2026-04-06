import { useEffect, useState, type FormEvent } from "react";
import "./App.css";

type Todo = {
  id: number;
  task: string;
  done: boolean;
};

const API_URL = "http://localhost:3000/todo";

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [text, setText] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const loadTodos = async () => {
    try {
      setLoading(true);
      setError("");

      const res = await fetch(API_URL);

      if (!res.ok) throw new Error();

      const data: Todo[] = await res.json();
      setTodos(data);
    } catch {
      setError("Cannot load todos");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadTodos();
  }, []);

  const handleAddTodo = async (e: FormEvent) => {
    e.preventDefault();

    const trimmed = text.trim();
    if (!trimmed) return;

    try {
      setError("");

      const res = await fetch(API_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          task: trimmed,
          done: false,
        }),
      });

      if (!res.ok) throw new Error();

      const newTodo: Todo = await res.json();

      setTodos((prev) => [...prev, newTodo]);
      setText("");
    } catch {
      setError("Cannot add todo");
    }
  };

  const handleToggle = async (todo: Todo) => {
    try {
      const res = await fetch(`${API_URL}/${todo.id}`, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          done: !todo.done,
        }),
      });

      if (!res.ok) throw new Error();

      setTodos((prev) =>
        prev.map((t) => (t.id === todo.id ? { ...t, done: !t.done } : t)),
      );
    } catch {
      setError("Cannot update todo");
    }
  };

  return (
    <div className="app">
      <div className="todo-card">
        <h1>My To Do App</h1>
        <p className="subtitle">React + Go Backend</p>

        <form onSubmit={handleAddTodo} className="todo-form">
          <input
            type="text"
            placeholder="Enter task..."
            value={text}
            onChange={(e) => setText(e.target.value)}
          />
          <button type="submit">Add</button>
        </form>

        {loading && <p className="message">Loading...</p>}
        {error && <p className="message error">{error}</p>}

        {!loading && todos.length === 0 && (
          <p className="message">No todos yet</p>
        )}

        <ul className="todo-list">
          {todos.map((todo) => (
            <li key={todo.id} className="todo-item">
              <label className="todo-row">
                <input
                  type="checkbox"
                  checked={todo.done}
                  onChange={() => handleToggle(todo)}
                />
                <span className={todo.done ? "done" : ""}>{todo.task}</span>
              </label>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default App;
