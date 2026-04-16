import type React from "react"
import type { JSX } from "react/jsx-dev-runtime"
import { Link } from "react-router-dom"

const EventsPage: React.FC = (): JSX.Element => {
    return (
        <main>
            <section className="section section--events">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">События</h2>
                            <p className="section__subtitle">Мастер-классы, встречи, лайв-кодинги и практикумы. Выберите удобные даты.</p>
                        </div>
                        <button className="btn btn--sm btn--outline">Сбросить</button>
                    </div>
                    <div className="events-grid">
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 25.04.2026 19:00–21:00</span>
                            </div>
                            <h3 className="event-card__title">Практикум: Введение в FastAPI</h3>
                            <p className="event-card__description">Быстрый старт с FastAPI: структура проекта, валидация схем и
                                тестирование.</p>
                            <p className="event-card__role">Наставник: Алексей Смирнов</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">10.05.2026 18:30–20:00</span>
                            </div>
                            <h3 className="event-card__title">Круглый стол: Архитектура микросервисов</h3>
                            <p className="event-card__description">Плюсы и минусы микросервисов, границы сервисов, коммуникации и
                                наблюдаемость.</p>
                            <p className="event-card__role">Наставники: Ольга Романова, Никита Лебедев</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 03.06.2026 20:00–21:30</span>
                            </div>
                            <h3 className="event-card__title">Лайв-кодинг: React + TypeScript</h3>
                            <p className="event-card__description">Пишем приложение с управлением состоянием, формами и оптимизациями
                                рендеринга.</p>
                            <p className="event-card__role">Наставник: Светлана Яковлева</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 22.05.2026 19:00–21:00</span>
                            </div>
                            <h3 className="event-card__title">Workshop: Profiling Go-сервисов</h3>
                            <p className="event-card__description">CPU/Memory профили, pprof, flamegraph и поиск узких мест в проде.</p>
                            <p className="event-card__role">Наставник: Екатерина Зайцева</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 30.04.2026 18:00–19:00</span>
                            </div>
                            <h3 className="event-card__title">Q&A с Solution Architect</h3>
                            <p className="event-card__description">Ответы на вопросы об архитектуре, интеграциях и безопасности приложений.
                            </p>
                            <p className="event-card__role">Наставник: Татьяна Соколова</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 12.07.2026 11:00–15:00</span>
                            </div>
                            <h3 className="event-card__title">Интенсив: Docker для разработчиков</h3>
                            <p className="event-card__description">Сборка образов, оптимизация Dockerfile, многоконтейнерные окружения и
                                best practices.</p>
                            <p className="event-card__role">Наставник: Игорь Кузнецов</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">05.05.2026 17:00–18:30</span>
                            </div>
                            <h3 className="event-card__title">AMA: Карьера в Data Science</h3>
                            <p className="event-card__description">Рынок, навыки, собесы и портфолио. Живые ответы на вопросы участников.
                            </p>
                            <p className="event-card__role">Наставник: Мария Орлова</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 18.06.2026 19:00–21:00</span>
                            </div>
                            <h3 className="event-card__title">Практикум: CI/CD с GitHub Actions</h3>
                            <p className="event-card__description">Сборка, тесты, релизы и деплой. Секреты, матрицы, кеширование и
                                артефакты.</p>
                            <p className="event-card__role">Наставник: Денис Пак</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 15.04.2026 18:00–19:00</span>
                            </div>
                            <h3 className="event-card__title">Разбор резюме и портфолио</h3>
                            <p className="event-card__description">Как структурировать опыт, выбрать проекты и выделиться в откликах.</p>
                            <p className="event-card__role">Ведущий: Команда «Студент и Т»</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </main>
    )
}

export default EventsPage