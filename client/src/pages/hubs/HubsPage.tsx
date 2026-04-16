import { useState } from "react"
import type React from "react"
import type { ChangeEvent, JSX } from "react"
import { useSelector } from "react-redux"
import { Link } from "react-router-dom"
import type { RootState } from "../../redux/store"
import { selectHubs } from "../../redux/features/hubs/hubSlice"

const HubsPage: React.FC = (): JSX.Element => {
    const [query, setQuery] = useState<string>("")
    const hubs = useSelector((state: RootState) => selectHubs(state))

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
                                <p className="hub-card__description">{hub.desription}</p>
                                <p className="hub-card__meta">📍 {hub.address}</p>
                                <p className="hub-card__meta">🕒 Сегодня: {hub.schedule}</p>

                                <div className="hub-card__actions">
                                    <Link to={`/hubs/${hub.id}`} className="btn btn--primary btn--sm">Открыть страницу хаба</Link>
                                </div>
                            </article>
                        ))}
                    </div>

                    {hubs.length === 0 ? (
                        <p className="hubs-search__empty">Хабы не найдены</p>
                    ) : null}
                </div>
            </section>
        </main>
    )
}

export default HubsPage
