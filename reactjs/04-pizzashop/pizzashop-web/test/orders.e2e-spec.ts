import { expect, test } from '@playwright/test'

test('list orders', async ({ page }) => {
  await page.goto('/orders', { waitUntil: 'networkidle' })

  expect(
    page.getByRole('cell', { name: 'Customer 1', exact: true }),
  ).toBeVisible()
  expect(
    page.getByRole('cell', { name: 'Customer 10', exact: true }),
  ).toBeVisible()
})

test.describe('paginate orders', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/orders', { waitUntil: 'networkidle' })
  })

  test('goes to next page', async ({ page }) => {
    await page.getByRole('button', { name: 'Próxima página' }).click()

    expect(
      page.getByRole('cell', { name: 'Customer 11', exact: true }),
    ).toBeVisible()
    expect(
      page.getByRole('cell', { name: 'Customer 20', exact: true }),
    ).toBeVisible()
  })

  test('goes to last page', async ({ page }) => {
    await page.getByRole('button', { name: 'Última página' }).click()

    expect(
      page.getByRole('cell', { name: 'Customer 51', exact: true }),
    ).toBeVisible()
    expect(
      page.getByRole('cell', { name: 'Customer 60', exact: true }),
    ).toBeVisible()
  })

  test('goes to previous page from last page', async ({ page }) => {
    await page.getByRole('button', { name: 'Última página' }).click()
    await page.getByRole('button', { name: 'Página anterior' }).click()

    expect(
      page.getByRole('cell', { name: 'Customer 41', exact: true }),
    ).toBeVisible()
    expect(
      page.getByRole('cell', { name: 'Customer 50', exact: true }),
    ).toBeVisible()
  })

  test('returns to first page from last page', async ({ page }) => {
    await page.getByRole('button', { name: 'Última página' }).click()
    await page.getByRole('button', { name: 'Primeira página' }).click()

    expect(
      page.getByRole('cell', { name: 'Customer 1', exact: true }),
    ).toBeVisible()
    expect(
      page.getByRole('cell', { name: 'Customer 10', exact: true }),
    ).toBeVisible()
  })
})

test.describe('filter orders', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/orders', { waitUntil: 'networkidle' })
  })

  test('filters by order id', async ({ page }) => {
    await page.getByPlaceholder('ID do pedido').fill('order-11')
    await page.getByRole('button', { name: 'Filtrar resultados' }).click()
    await expect(page.getByRole('cell', { name: 'order-11' })).toBeVisible()
  })

  test('filters by customer name', async ({ page }) => {
    await page.getByPlaceholder('Nome do cliente').fill('Customer 11')

    await page.getByRole('button', { name: 'Filtrar resultados' }).click()

    await expect(page.getByRole('cell', { name: 'Customer 11' })).toBeVisible()
  })

  test('filters by status', async ({ page }) => {
    await page.getByRole('combobox').click()

    await page.getByLabel('Pendente').click()

    await page.getByRole('button', { name: 'Filtrar resultados' }).click()

    const tableRows = await page.getByRole('cell', { name: 'Pendente' })

    await expect(tableRows).toHaveCount(10)
  })
})
