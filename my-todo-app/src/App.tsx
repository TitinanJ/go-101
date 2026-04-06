import { useEffect, useState, type FormEvent } from "react";
import "./App.css";

type Todo = {
  id: number;
  text: string;
  done: boolean;
};

const API_URL = "http://localhost:8080/todos";

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [text, setText] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const loadTodos = async () => {
    try {
      setLoading(true);
      setError("");

      const response = await fetch(API_URL);

      if (!response.ok) {
        throw new Error("Failed to load todos");
      }

      const data: Todo[] = await response.json();
      setTodos(data);
    } catch (err) {
      setError("Cannot load todos");
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadTodos();
  }, []);

  const handleAddTodo = async (e: FormEvent) => {
    e.preventDefault();

    const trimmedText = text.trim();

    if (!trimmedText) return;

    try {
      setError("");

      const response = await fetch(API_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          text: trimmedText,
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to add todo");
      }

      const newTodo: Todo = await response.json();

      setTodos((prev) => [...prev, newTodo]);
      setText("");
    } catch (err) {
      setError("Cannot add todo");
      console.error(err);
    }
  };

  const handleDeleteTodo = async (id: number) => {
    try {
      setError("");

      const response = await fetch(`${API_URL}/${id}`, {
        method: "DELETE",
      });

      if (!response.ok) {
        throw new Error("Failed to delete todo");
      }

      setTodos((prev) => prev.filter((todo) => todo.id !== id));
    } catch (err) {
      setError("Cannot delete todo");
      console.error(err);
    }
  };

  return (
    <div className="app">
      <div className="todo-card">
        <h1>My To Do App</h1>
        <p className="subtitle">React Frontend for Go Backend</p>

        <form onSubmit={handleAddTodo} className="todo-form">
          <input
            type="text"
            placeholder="Enter a new task"
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
              <span className={todo.done ? "done" : ""}>{todo.text}</span>
              <button onClick={() => handleDeleteTodo(todo.id)}>Delete</button>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default App;
