import { zodResolver } from '@hookform/resolvers/zod'
import { Search, X } from 'lucide-react'
import { Controller, useForm } from 'react-hook-form'
import { useSearchParams } from 'react-router'
import z4 from 'zod/v4'

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'

const orderFilterSchema = z4.object({
  orderId: z4.string().optional(),
  customerName: z4.string().optional(),
  status: z4.string().optional(),
})

type OrderFilterSchema = z4.infer<typeof orderFilterSchema>

export function OrderTableFilters() {
  const [searchParams, setSearchParams] = useSearchParams()
  const orderId = searchParams.get('orderId')
  const customerName = searchParams.get('customerName')
  const status = searchParams.get('status')

  const { register, handleSubmit, control, reset } = useForm<OrderFilterSchema>(
    {
      resolver: zodResolver(orderFilterSchema),
      defaultValues: {
        orderId: orderId ?? '',
        customerName: customerName ?? '',
        status: status ?? ' all',
      },
    },
  )

  function handleFilter({ orderId, customerName, status }: OrderFilterSchema) {
    setSearchParams((state) => {
      if (orderId) {
        state.set('orderId', orderId)
      } else {
        state.delete('orderId')
      }
      if (customerName) {
        state.set('customerName', customerName)
      } else {
        state.delete('customerName')
      }
      if (status) {
        state.set('status', status)
      } else {
        state.delete('status')
      }

      state.set('page', '1')
      return state
    })
  }

  function handleClearFilters() {
    setSearchParams((state) => {
      state.delete('orderId')
      state.delete('customerName')
      state.delete('status')
      state.set('page', '1')

      return state
    })

    reset({
      orderId: '',
      customerName: '',
      status: 'all',
    })
  }

  return (
    <form
      onSubmit={handleSubmit(handleFilter)}
      className="flex items-center gap-2"
    >
      <span className="text-sm font-semibold">Filtros</span>
      <Input
        placeholder="ID do pedido"
        className="h-8 w-auto"
        {...register('orderId')}
      />
      <Input
        placeholder="Nome do cliente"
        className="h-8 w-[320px]"
        {...register('customerName')}
      />
      <Controller
        name="status"
        control={control}
        render={({ field: { name, onChange, value, disabled } }) => {
          return (
            <Select
              defaultValue="all"
              name={name}
              onValueChange={onChange}
              value={value}
              disabled={disabled}
            >
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
          )
        }}
      />

      <Button type="submit" variant="secondary" size="xs">
        <Search className="mr-2 h-4 w-4" /> Filtrar resultados
      </Button>
      <Button
        type="button"
        variant="outline"
        size="xs"
        onClick={handleClearFilters}
      >
        <X className="mr-2 h-4 w-4" /> Remover Filtros
      </Button>
    </form>
  )
}
