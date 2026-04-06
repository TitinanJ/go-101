export type Todo = {
  id: number;
  task: string;
  done: boolean;
};

const API_URL = "http://localhost:3000/todo";

export async function getTodos(): Promise<Todo[]> {
    const response = await fetch(API_URL);

    if (!response.ok) {
        throw new Error("Failed to load todos");
    }

    return response.json();
}

export async function createTodo(task: string): Promise<Todo> {
    const response = await fetch(API_URL, {
        method: "POST",
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify({
        task,
        done: false,
        }),
    });

    if (!response.ok) {
        throw new Error("Failed to create todo");
    }

    return response.json();
}

export async function updateTodoDone(id: number, done: boolean): Promise<void> {
    const response = await fetch(`${API_URL}/${id}`, {
        method: "PATCH",
        headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify({ done }),
    });

    if (!response.ok) {
        throw new Error("Failed to update todo");
    }
}
