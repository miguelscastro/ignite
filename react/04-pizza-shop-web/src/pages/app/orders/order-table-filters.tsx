import { Search, X } from "lucide-react";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

export function OrderTableFilters() {
  return (
    <form className="flex items-center gap-2">
      <span className="text-sm font-semibold">Filtros</span>
      <Input placeholder="ID do pedido" className="h-8 w-auto" />
      <Input placeholder="Nome do cliente" className="h-8 w-[320px]" />
      <Select defaultValue="all">
        <SelectTrigger className="h-8 w-[180px] cursor-pointer">
          <SelectValue />
        </SelectTrigger>
        <SelectContent>
          <SelectItem className="cursor-pointer" value="all">
            Todos status
          </SelectItem>
          <SelectItem className="cursor-pointer" value="pending">
            Pendente
          </SelectItem>
          <SelectItem className="cursor-pointer" value="canceled">
            Cancelado
          </SelectItem>
          <SelectItem className="cursor-pointer" value="processing">
            Em preparo
          </SelectItem>
          <SelectItem className="cursor-pointer" value="delivering">
            Em entrega
          </SelectItem>
          <SelectItem className="cursor-pointer" value="delivered">
            Entregue
          </SelectItem>
        </SelectContent>
      </Select>

      <Button type="submit" variant="secondary" size="xs">
        <Search className="mr-2 h-4 w-4" /> Filtrar resultados
      </Button>
      <Button type="button" variant="outline" size="xs">
        <X className="mr-2 h-4 w-4" /> Filtrar resultados
      </Button>
    </form>
  );
}
