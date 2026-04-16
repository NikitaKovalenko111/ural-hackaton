import { useState } from "react"
import type { JSX } from "react"
import type React from "react"
import { useSelector } from "react-redux"
import { Link, Navigate } from "react-router-dom"
import { useParams } from "react-router-dom"
import RegistrationModal from "../../components/registrationModal/RegistrationModal"
import type { RootState } from "../../redux/store"
import { selectHubs } from "../../redux/features/hubs/hubSlice"

const IndexPage: React.FC = (): JSX.Element => {
    const { hubId } = useParams<{ hubId: string }>()
    const hubs = useSelector((state: RootState) => selectHubs(state))
    const [isRegistrationOpen, setIsRegistrationOpen] = useState<boolean>(false)
    const [selectedEventTitle, setSelectedEventTitle] = useState<string>("Ближайшее событие хаба")
    const hub = hubs.find((hubItem) => hubItem.id === parseInt(hubId as string)) ?? hubs[0]

    if (!hub) {
        return <Navigate to="/hubs" replace />
    }

    const openRegistrationModal = (eventTitle: string): void => {
        setSelectedEventTitle(eventTitle)
        setIsRegistrationOpen(true)
    }

    const statusText = hub.status === "open" ? "● Открыт" : hub.status === "busy" ? "● Высокая загрузка" : "● Закрыт"

    return (
        <main>
            <section className="hero">
                <div className="container hero__card">
                    <div className="hero__info">
                        <div className="status-header">
                            <div className="status-header__icon">🕒</div>
                            <div>
                                <h1 className="status-header__title">{hub.name}</h1>
                                <p className="status-header__schedule">График сегодня: {hub.schedule}</p>
                            </div>
                            <span className="badge--open">{statusText}</span>
                        </div>
                        <div className="stats__container">
                            <div className="stats__box">
                                <div className="stats__label">Количество людей сейчас</div>
                                <div className="stats__value">{hub.occupancy}</div>
                                <div className="stats__unit">чел.</div>
                                <div className="stats__update-text">Обн. 5 мин</div>
                            </div>
                            <div className="event-box">
                                <div className="event-box__header">
                                    <span>Ближайшее событие</span>
                                    <Link to="/events" className="link--small">Все события</Link>
                                </div>
                                {/*<div className="event-box__details">
                                    <h3>{hub.nearestEvent.title}</h3>
                                    <p>Начало: {hub.nearestEvent.time}</p>
                                </div>*/}
                                <button
                                    type="button"
                                    className="btn btn--primary"
                                    onClick={() => openRegistrationModal("Ближайшее событие хаба")}
                                >
                                    Записаться
                                </button>
                                <div className="event-box__status">🕒 Сейчас {hub.schedule} → {hub.status === "closed" ? "закрыт" : "открыт"}</div>
                            </div>
                        </div>
                    </div>
                    <div className="hero__image-wrapper">
                        <div className="hero__image"></div>
                        <div className="image-overlay">description</div>
                        <div className="wifi-badge">📶 Wi-Fi: free</div>
                    </div>
                </div>
            </section>

            <section className="section section--hubs-overview">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">О хабе</h2>
                            <p className="section__subtitle">{hub.desription}</p>
                        </div>
                        <Link to="/hubs" className="btn btn--outline btn--sm">Все хабы</Link>
                    </div>
                    <div className="contacts__wrapper">
                        <div className="contacts__info">
                            <div className="contacts__item">📍 {hub.address}</div>
                            {/*<div className="contacts__item">📞 {hub.contacts.phone}</div>
                            <div className="contacts__item">✉️ {hub.contacts.email}</div>
                            <div className="contacts__item">✈️ {hub.contacts.telegram}</div>*/}
                        </div>
                        <div className="contacts__map">
                            <div className="map-placeholder"></div>
                        </div>
                    </div>
                </div>
            </section>

            <section className="section section--events">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">Ближайшие события</h2>
                            <p className="section__subtitle">Следите за расписанием и присоединяйтесь</p>
                        </div>
                        <Link to="/events" className="btn btn--primary btn--sm">Все события</Link>
                    </div>
                    <div className="events-grid">
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 18:00</span>
                                <span className="event-card__date">сегодня</span>
                            </div>
                            <h3 className="event-card__title">Нетворкинг: Проекты сообщества</h3>
                            <p className="event-card__description">Открытый круг знакомств и обмена идеями.</p>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 19:00</span>
                                <span className="event-card__date">сегодня</span>
                            </div>
                            <h3 className="event-card__title">Мастер-класс: Веб-дизайн</h3>
                            <p className="event-card__description">Практика UI-решений и прототипирование.</p>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 20:30</span>
                                <span className="event-card__date">сегодня</span>
                            </div>
                            <h3 className="event-card__title">Кинопоказ + обсуждение</h3>
                            <p className="event-card__description">Вдохновляющее кино и разговор после.</p>
                        </div>
                    </div>
                </div>
            </section>

            <section className="section section--mentors">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">Лучшие менторы</h2>
                            <p className="section__subtitle">Эксперты, которые помогут ускорить ваш рост</p>
                        </div>
                        <Link to="/mentors" className="btn btn--primary btn--sm">Стать ментором</Link>
                    </div>
                    <div className="mentors-grid">
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">👤 UI/UX</span>
                                <span className="rating">рейтинг 4.9</span>
                            </div>
                            <h3 className="mentor-card__title">Александр Петров</h3>
                            <p className="mentor-card__role">Продуктовый дизайнер • 8 лет опыта</p>
                            <div className="mentor-card__actions">
                                <Link to="/mentors" className="btn btn--primary btn--sm">Подробнее</Link>
                                <Link to="/mentors" className="link--text">Портфолио</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">&lt;/&gt; Frontend</span>
                                <span className="rating">рейтинг 4.8</span>
                            </div>
                            <h3 className="mentor-card__title">Мария Орлова</h3>
                            <p className="mentor-card__role">Senior Frontend • React, TypeScript</p>
                            <div className="mentor-card__actions">
                                <Link to="/mentors" className="btn btn--primary btn--sm">Подробнее</Link>
                                <Link to="/mentors" className="link--text">Расписание</Link>
                            </div>
                        </div>
                        <div className="mentor-card">
                            <div className="mentor-card__top">
                                <span className="tag">📊 Data</span>
                                <span className="rating">рейтинг 5.0</span>
                            </div>
                            <h3 className="mentor-card__title">Илья Громов</h3>
                            <p className="mentor-card__role">ML/Analytics • Python, SQL, DS</p>
                            <div className="mentor-card__actions">
                                <Link to="/mentors" className="btn btn--primary btn--sm">Подробнее</Link>
                                <Link to="/mentors" className="link--text">Консультация</Link>
                            </div>
                        </div>
                    </div>
                </div>
            </section>

            <section className="section section--contacts">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">Контакты</h2>
                            <p className="section__subtitle">Всегда на связи — приходите, звоните, пишите</p>
                        </div>
                        <Link to="/events" className="btn btn--primary btn--sm">Как добраться</Link>
                    </div>
                    <div className="contacts__wrapper">
                        <div className="contacts__info">
                            <div className="contacts__item">📍 Город, Ул. Примерная, 15</div>
                            <div className="contacts__item">📞 +7 (000) 000-00-00</div>
                            <div className="contacts__item">✉️ hello@yellowhub.example</div>
                            <div className="contacts__item">✈️ Telegram</div>
                        </div>
                        <div className="contacts__map">
                            <div className="map-placeholder"></div>
                        </div>
                    </div>
                </div>
            </section>

            <RegistrationModal
                isOpen={isRegistrationOpen}
                onClose={() => setIsRegistrationOpen(false)}
                eventTitle={selectedEventTitle}
            />
        </main>
    )
}

export default IndexPage