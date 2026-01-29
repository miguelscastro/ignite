import { render } from '@testing-library/react'
import { userEvent } from '@testing-library/user-event'

import { Pagination } from './pagination'

const onPageChangeCallback = vi.fn()

// vi === acesso aos utils do vitest
/*
	fn === um spy de função que apenas recebe os parâmetros que seriam 
	passados para a função real, sem executar nenhuma ação
*/

describe('Pagination', () => {
  beforeEach(() => {
    onPageChangeCallback.mockClear()
  })

  it('should display the right amount of pages and results', () => {
    const wrapper = render(
      <Pagination
        pageIndex={0}
        totalCount={200}
        perPage={10}
        onPageChange={() => {}}
      />,
    )

    expect(wrapper.getByText('Página 1 de 20')).toBeInTheDocument()
    expect(wrapper.getByText('Total de 200 item(s)')).toBeInTheDocument()
  })
  it('should be able to navigate to the next page', async () => {
    const user = userEvent.setup()

    const wrapper = render(
      <Pagination
        pageIndex={0}
        totalCount={200}
        perPage={10}
        onPageChange={onPageChangeCallback}
      />,
    )

    const nextPageButton = wrapper.getByRole('button', {
      name: 'Próxima página',
    })

    await user.click(nextPageButton)

    expect(onPageChangeCallback).toHaveBeenCalledWith(1)
  })
  it('should be able to navigate to the previous page', async () => {
    const user = userEvent.setup()

    const wrapper = render(
      <Pagination
        pageIndex={5}
        totalCount={200}
        perPage={10}
        onPageChange={onPageChangeCallback}
      />,
    )

    const prevPageButton = wrapper.getByRole('button', {
      name: 'Página anterior',
    })

    await user.click(prevPageButton)

    expect(onPageChangeCallback).toHaveBeenCalledWith(4)
  })
  it('should be able to navigate to the first page', async () => {
    const user = userEvent.setup()

    const wrapper = render(
      <Pagination
        pageIndex={5}
        totalCount={200}
        perPage={10}
        onPageChange={onPageChangeCallback}
      />,
    )

    const firstPageButton = wrapper.getByRole('button', {
      name: 'Primeira página',
    })

    await user.click(firstPageButton)

    expect(onPageChangeCallback).toHaveBeenCalledWith(0)
  })
  it('should be able to navigate to the last page', async () => {
    const user = userEvent.setup()

    const wrapper = render(
      <Pagination
        pageIndex={5}
        totalCount={200}
        perPage={10}
        onPageChange={onPageChangeCallback}
      />,
    )

    const lastPageButton = wrapper.getByRole('button', {
      name: 'Última página',
    })

    await user.click(lastPageButton)

    expect(onPageChangeCallback).toHaveBeenCalledWith(19)
  })
})
