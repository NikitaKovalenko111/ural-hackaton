import { useEffect, useState } from "react"
import type React from "react"
import type { ChangeEvent, JSX } from "react"
import { isAxiosError } from "axios"
import { Link } from "react-router-dom"
import { hubsApi } from "../../api/api"
import type { IHub } from "../../types"

const HubsPage: React.FC = (): JSX.Element => {
    const [query, setQuery] = useState<string>("")
    const [hubs, setHubs] = useState<IHub[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [errorMessage, setErrorMessage] = useState<string>("")

    useEffect(() => {
        const loadHubs = async (): Promise<void> => {
            setIsLoading(true)
            setErrorMessage("")

            try {
                const normalizedQuery = query.trim()
                const data = normalizedQuery
                    ? await hubsApi.searchHubs(normalizedQuery)
                    : await hubsApi.getHubs()

                setHubs(data)
            } catch (error) {
                if (isAxiosError(error) && typeof error.response?.data?.message === "string") {
                    setErrorMessage(error.response.data.message)
                } else {
                    setErrorMessage("Не удалось загрузить хабы")
                }
                setHubs([])
            } finally {
                setIsLoading(false)
            }
        }

        void loadHubs()
    }, [query])

    const handleQueryChange = (event: ChangeEvent<HTMLInputElement>): void => {
        setQuery(event.target.value)
    }

    return (
        <main>
            <section className="section section--hubs">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h1 className="section__title">Хабы</h1>
                            <p className="section__subtitle">Найдите ближайший хаб и перейдите на его страницу с расписанием и статусом.</p>
                        </div>
                    </div>

                    <div className="hubs-search">
                        <label className="hubs-search__label" htmlFor="hubs-search-input">Поиск по городу, названию или адресу</label>
                        <input
                            id="hubs-search-input"
                            type="search"
                            value={query}
                            onChange={handleQueryChange}
                            placeholder="Например: Екатеринбург"
                        />
                    </div>

                    <div className="hubs-grid">
                        {hubs.map((hub) => (
                            <article className="hub-card" key={hub.id}>
                                <div className="hub-card__header">
                                    <span className={`hub-card__status hub-card__status--${hub.status}`}>
                                        {hub.status === "open" ? "Открыт" : hub.status === "busy" ? "Высокая загрузка" : "Закрыт"}
                                    </span>
                                    <span className="hub-card__city">{hub.city}</span>
                                </div>

                                <h3 className="hub-card__title">{hub.name}</h3>
                                <p className="hub-card__description">{hub.description}</p>
                                <p className="hub-card__meta">📍 {hub.address}</p>
                                <p className="hub-card__meta">🕒 Сегодня: {hub.schedule}</p>

                                <div className="hub-card__actions">
                                    <Link to={`/hubs/${hub.id}`} className="btn btn--primary btn--sm">Открыть страницу хаба</Link>
                                </div>
                            </article>
                        ))}
                    </div>

                    {isLoading ? (
                        <p className="hubs-search__empty">Загрузка хабов...</p>
                    ) : null}

                    {errorMessage ? (
                        <p className="hubs-search__empty">{errorMessage}</p>
                    ) : null}

                    {!isLoading && !errorMessage && hubs.length === 0 ? (
                        <p className="hubs-search__empty">
                            {query.trim() ? "По вашему запросу хабы не найдены" : "Хабы не найдены"}
                        </p>
                    ) : null}
                </div>
            </section>
        </main>
    )
}

export default HubsPage
