import { useQuery } from '@tanstack/react-query'
import { subDays } from 'date-fns'
import { useMemo, useState } from 'react'
import { type DateRange } from 'react-day-picker'
import {
  CartesianGrid,
  Line,
  LineChart,
  ResponsiveContainer,
  XAxis,
  YAxis,
} from 'recharts'
import colors from 'tailwindcss/colors'

import { getDailyRevenueInPeriod } from '@/api/get-daily-revenue-in-period'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { DateRangePicker } from '@/components/ui/date-range-picker'
import { Label } from '@/components/ui/label'

export function RevenueChart() {
  const [dateRangeSample, setDateRangeSample] = useState<DateRange | undefined>(
    {
      from: subDays(new Date(), 7),
      to: new Date(),
    },
  )

  const { data: dailyRevenueInPeriod } = useQuery({
    queryKey: ['metrics, daily-revenue-in-period', dateRangeSample],
    queryFn: () =>
      getDailyRevenueInPeriod({
        from: dateRangeSample?.from,
        to: dateRangeSample?.to,
      }),
  })

  const chartDate = useMemo(() => {
    return dailyRevenueInPeriod?.map((chartItem) => {
      return {
        date: chartItem.date,
        receipt: chartItem.receipt / 100,
      }
    })
  }, [dailyRevenueInPeriod])

  return (
    <Card className="col-span-6">
      <CardHeader className="ml-2.5 flex-row items-start justify-between pb-8">
        <div className="space-y-1">
          <CardTitle className="text-base font-medium">
            Receita no período
          </CardTitle>
          <CardDescription>Receita diária no perído</CardDescription>
        </div>
        <div className="flex items-center gap-3">
          <Label>Período</Label>
          <DateRangePicker
            date={dateRangeSample}
            onDateChange={setDateRangeSample}
          />
        </div>
      </CardHeader>
      <CardContent>
        {dailyRevenueInPeriod && (
          <ResponsiveContainer width="100%" height={240}>
            <LineChart data={chartDate} style={{ fontSize: 12 }}>
              <XAxis dataKey="date" tickLine={false} axisLine={false} dy={16} />

              <YAxis
                stroke="#888"
                axisLine={false}
                tickLine={false}
                width={80}
                tickFormatter={(value: number) =>
                  value.toLocaleString('pt-BR', {
                    style: 'currency',
                    currency: 'BRL',
                  })
                }
              />
              <CartesianGrid vertical={false} className="stroke-muted" />

              <Line
                type="linear"
                strokeWidth={2}
                dataKey="receipt"
                stroke={colors.violet[500]}
              />
            </LineChart>
          </ResponsiveContainer>
        )}
      </CardContent>
    </Card>
  )
}
