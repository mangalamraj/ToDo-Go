"use client";
import { useState } from "react";
import { Button } from "./ui/button";
import { Input } from "./ui/input";
import TodoComponent from "./todoComponent";

const TodoForm = () => {
  const [title, setTitle] = useState("");
  const addTodo = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await fetch("http://localhost:8080/addtodo", {
        headers: { "Content-Type": "application/json" },
        method: "POST",
        body: JSON.stringify({ title }),
      });
      if (response.ok) {
        console.log("Todo updated!");
      }
    } catch (e) {
      console.log("Error adding todo", e);
    }
  };
  return (
    <div className="flex flex-col justify-center align-middle items-center h-screen gap-2">
      <div className="container flex justify-center ">
        <form onSubmit={addTodo} className="flex m-auto w-[600px] gap-4">
          <Input
            value={title}
            onChange={(e) => {
              setTitle(e.target.value);
            }}
          />
          <Button type="submit">Add</Button>
        </form>
      </div>
      <TodoComponent />
    </div>
  );
};
export default TodoForm;
