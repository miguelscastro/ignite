import { Link, useRouteError } from 'react-router'

export function Error() {
  const error = useRouteError() as Error

  return (
    <div className="flex h-screen flex-col items-center justify-center gap-2">
      <h1 className="text-4xl font-bold">Whoops, algo aconteceu...</h1>
      <p className="text-accent-foreground mb-15">
        Um erro aconteceu na aplicação, abaixo você encontra mais detalhes:
      </p>
      <div className="bg-card max-h-82 w-150 max-w-4xl overflow-y-auto rounded-2xl p-3 [&::-webkit-scrollbar]:w-2 [&::-webkit-scrollbar-thumb]:rounded-full [&::-webkit-scrollbar-thumb]:bg-gray-400 dark:[&::-webkit-scrollbar-thumb]:bg-gray-600 [&::-webkit-scrollbar-track]:bg-gray-100 dark:[&::-webkit-scrollbar-track]:bg-gray-800">
        <pre className="text-left break-words whitespace-pre-wrap text-red-500">
          {error?.message || JSON.stringify(error, null, 2)}
        </pre>
      </div>
      <p className="text-accent-foreground mt-10">
        Voltar para o{' '}
        <Link to="/" className="text-sky-500 dark:text-sky-400">
          Dashboard
        </Link>
      </p>
    </div>
  )
}
