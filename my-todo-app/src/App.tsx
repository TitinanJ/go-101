import { useEffect, useState, type FormEvent } from "react";

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
      setError("");

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
    <div className="min-h-screen bg-stone-100 px-6 py-10 flex items-center justify-center">
      <div className="w-full max-w-lg rounded-3xl border border-black/10 bg-stone-50 p-8 shadow-[0_2px_0_rgba(0,0,0,0.06),0_12px_40px_rgba(0,0,0,0.07)]">
        <h1 className="text-4xl font-serif text-stone-900 tracking-tight">
          My To Do App
        </h1>

        <p className="mt-1 mb-7 text-[11px] uppercase tracking-[0.18em] text-stone-400">
          React + Go Backend
        </p>

        <form onSubmit={handleAddTodo} className="mb-6 flex gap-2">
          <input
            type="text"
            placeholder="Enter task..."
            value={text}
            onChange={(e) => setText(e.target.value)}
            className="flex-1 rounded-xl border border-black/10 bg-white px-4 py-2.5 text-sm text-stone-900 outline-none placeholder:text-stone-300 focus:border-amber-700 focus:ring-4 focus:ring-amber-700/10"
          />
          <button
            type="submit"
            className="rounded-xl bg-stone-900 px-5 py-2.5 text-sm font-medium text-stone-50 transition hover:bg-stone-800 active:scale-95"
          >
            Add
          </button>
        </form>

        {loading && (
          <p className="py-4 text-center text-sm text-stone-400">Loading...</p>
        )}

        {error && (
          <p className="rounded-lg bg-red-50 px-4 py-3 text-center text-sm text-red-500">
            {error}
          </p>
        )}

        {!loading && todos.length === 0 && (
          <p className="py-4 text-center text-sm text-stone-400">
            No todos yet
          </p>
        )}

        <ul className="mt-2 list-none p-0">
          {todos.map((todo) => (
            <li
              key={todo.id}
              className="border-b border-black/5 py-3 last:border-b-0"
            >
              <label className="flex w-full cursor-pointer items-center gap-3">
                <input
                  type="checkbox"
                  checked={todo.done}
                  onChange={() => handleToggle(todo)}
                  className="h-4 w-4 cursor-pointer rounded border-stone-300 accent-stone-900"
                />
                <span
                  className={`flex-1 wrap-break-words text-sm ${
                    todo.done ? "text-stone-300" : "text-stone-900"
                  }`}
                >
                  {todo.task}
                </span>
              </label>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default App;
