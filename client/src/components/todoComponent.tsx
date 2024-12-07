import { Trash2 } from "lucide-react";

const TodoComponent = () => {
  return (
    <div className="border-2 p-2 rounded-md w-[600px] justify-between flex items-center">
      <div>hi</div>
      <Trash2 color="red" size={18} />
    </div>
  );
};

export default TodoComponent;
