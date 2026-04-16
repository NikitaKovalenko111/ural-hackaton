import type { JSX } from "react"
import type React from "react"
import { Link } from "react-router-dom"

const NotFound: React.FC = (): JSX.Element => {
    return (
        <main>
            <section className="error-layout">
                <div className="container">
                    <div className="error-card">
                        <div className="error-icon">🔍</div>
                        <div className="error-code">404</div>
                        <h1 className="error-title">Страница не найдена</h1>
                        <p className="error-message">
                            Похоже, вы забрели не туда. Страница, которую вы ищете,
                            была перемещена, удалена или никогда не существовала.
                        </p>

                        <div className="error-actions">
                            <Link to="/hubs" className="btn btn--primary">На страницу хабов</Link>
                            <Link to="/events" className="btn btn--outline">События</Link>
                        </div>

                        <div className="help-links">
                            <h3>Возможно, вы искали:</h3>
                            <div className="help-grid">
                                <Link to="/hubs" className="help-link">
                                    <span>🏠</span> Хабы
                                </Link>
                                <Link to="/events" className="help-link">
                                    <span>📅</span> Ближайшие события
                                </Link>
                                <Link to="/mentors" className="help-link">
                                    <span>👨‍🏫</span> Найти ментора
                                </Link>
                                <Link to="/profile" className="help-link">
                                    <span>👤</span> Профиль
                                </Link>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </main>
    )
}

export default NotFound