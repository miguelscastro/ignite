import { test, expect } from '@playwright/test'

test('update profile succesfully', async ({ page }) => {
  await page.goto('/', { waitUntil: 'networkidle' })
  await page.getByRole('button', { name: 'Pizza Shop' }).click()
  await page.getByRole('menuitem', { name: 'Perfil da loja' }).click()
  await page.getByLabel('Nome').fill('Papa pizzaria')
  await page.getByLabel('Descrição').fill('Test Description')
  await page.getByRole('button', { name: 'Salvar' }).click()
  await page.waitForLoadState('networkidle')
  const toast = page.getByText('Perfil atualizado com sucesso')
  expect(toast).toBeVisible()
  await page.getByRole('button', { name: 'Close' }).click()
  await page.waitForTimeout(250)
  expect(page.getByRole('button', { name: 'Papa pizzaria' })).toBeVisible()
})
