import { isAxiosError } from "axios"
import { useEffect, useMemo, useState } from "react"
import type React from "react"
import type { ChangeEvent, FormEvent, JSX } from "react"
import { useSelector } from "react-redux"
import { eventsApi } from "../../api/api"
import RegistrationModal from "../../components/registrationModal/RegistrationModal"
import { selectUser } from "../../redux/features/users/userSlice"
import type { IEvent } from "../../types"

type EventFormState = {
    name: string
    description: string
    startTime: string
    endTime: string
    hubId: string
    mentorId: string
}

const initialEventForm: EventFormState = {
    name: "",
    description: "",
    startTime: "",
    endTime: "",
    hubId: "",
    mentorId: "",
}

const formatEventDateTime = (value: string): string => {
    const date = new Date(value)
    if (Number.isNaN(date.getTime())) {
        return value
    }

    return new Intl.DateTimeFormat("ru-RU", {
        day: "2-digit",
        month: "2-digit",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit",
    }).format(date)
}

const EventsPage: React.FC = (): JSX.Element => {
    const user = useSelector(selectUser)
    const isMentor = user?.role?.toLowerCase() === "mentor"
    const isAdmin = user?.role?.toLowerCase() === "admin"
    const canCreateEvent = isMentor || isAdmin

    const [query, setQuery] = useState<string>("")
    const [events, setEvents] = useState<IEvent[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(true)
    const [errorMessage, setErrorMessage] = useState<string>("")
    const [isRegistrationOpen, setIsRegistrationOpen] = useState<boolean>(false)
    const [selectedEventTitle, setSelectedEventTitle] = useState<string>("Выбранное событие")

    const [eventForm, setEventForm] = useState<EventFormState>(initialEventForm)
    const [eventSubmitMessage, setEventSubmitMessage] = useState<string>("")
    const [eventSubmitError, setEventSubmitError] = useState<string>("")
    const [isSubmittingEvent, setIsSubmittingEvent] = useState<boolean>(false)

    const normalizedQuery = useMemo(() => query.trim(), [query])

    const loadEvents = async (searchValue: string): Promise<void> => {
        setIsLoading(true)
        setErrorMessage("")

        try {
            const data = searchValue
                ? await eventsApi.searchEvents(searchValue)
                : await eventsApi.getEvents()
            setEvents(data)
        } catch (error) {
            setEvents([])
            if (isAxiosError(error) && typeof error.response?.data?.message === "string") {
                setErrorMessage(error.response.data.message)
            } else {
                setErrorMessage("Не удалось загрузить события")
            }
        } finally {
            setIsLoading(false)
        }
    }

    useEffect(() => {
        void loadEvents(normalizedQuery)
    }, [normalizedQuery])

    const handleQueryChange = (event: ChangeEvent<HTMLInputElement>): void => {
        setQuery(event.target.value)
    }

    const handleReset = (): void => {
        setQuery("")
    }

    const handleOpenRegistration = (eventTitle: string): void => {
        setSelectedEventTitle(eventTitle)
        setIsRegistrationOpen(true)
    }

    const handleCreateEvent = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault()
        setEventSubmitMessage("")
        setEventSubmitError("")
        setIsSubmittingEvent(true)

        try {
            const startIso = new Date(eventForm.startTime).toISOString()
            const endIso = new Date(eventForm.endTime).toISOString()

            await eventsApi.saveEvent({
                name: eventForm.name.trim(),
                description: eventForm.description.trim(),
                start_time: startIso,
                end_time: endIso,
                hub_id: Number(eventForm.hubId),
                mentor_id: eventForm.mentorId ? Number(eventForm.mentorId) : undefined,
            })

            setEventForm(initialEventForm)
            setEventSubmitMessage("Событие создано")
            await loadEvents(normalizedQuery)
        } catch (error) {
            if (isAxiosError(error) && typeof error.response?.data?.message === "string") {
                setEventSubmitError(error.response.data.message)
            } else {
                setEventSubmitError("Не удалось создать событие")
            }
        } finally {
            setIsSubmittingEvent(false)
        }
    }

    return (
        <main>
            <section className="section section--events">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">События</h2>
                            <p className="section__subtitle">События загружаются с сервера. Используйте поиск по названию.</p>
                        </div>
                        <button type="button" className="btn btn--sm btn--outline" onClick={handleReset}>Сбросить</button>
                    </div>

                    <div className="hubs-search">
                        <label className="hubs-search__label" htmlFor="events-search-input">Поиск событий по названию</label>
                        <input
                            id="events-search-input"
                            type="search"
                            value={query}
                            onChange={handleQueryChange}
                            placeholder="Например: FastAPI"
                        />
                    </div>

                    <div className="events-grid">
                        {events.map((eventItem) => (
                            <article className="event-card" key={eventItem.id}>
                                <div className="event-card__header">
                                    <span className="event-card__time">🕒 {formatEventDateTime(eventItem.start)} - {formatEventDateTime(eventItem.end)}</span>
                                </div>
                                <h3 className="event-card__title">{eventItem.name}</h3>
                                <p className="event-card__description">{eventItem.description}</p>
                                <p className="event-card__role">Хаб ID: {eventItem.hubId}</p>
                                <p className="event-card__role">Ментор ID: {eventItem.mentorId ?? "не указан"}</p>
                                <div className="mentor-card__actions">
                                    <button type="button" className="btn btn--primary btn--sm" onClick={() => handleOpenRegistration(eventItem.name)}>
                                        Регистрация
                                    </button>
                                </div>
                            </article>
                        ))}
                    </div>

                    {isLoading ? <p className="hubs-search__empty">Загрузка событий...</p> : null}
                    {errorMessage ? <p className="hubs-search__empty">{errorMessage}</p> : null}
                    {!isLoading && !errorMessage && events.length === 0 ? (
                        <p className="hubs-search__empty">{normalizedQuery ? "События по запросу не найдены" : "События не найдены"}</p>
                    ) : null}
                </div>
            </section>

            {canCreateEvent ? (
                <section className="section section--contacts">
                    <div className="container">
                        <div className="section__header">
                            <div>
                                <h2 className="section__title">Создание события</h2>
                                <p className="section__subtitle">Форма доступна ментору и администратору.</p>
                            </div>
                        </div>

                        <article className="booking-card booking-card--event-create">
                            <form className="booking-form booking-form--event-create" onSubmit={handleCreateEvent}>
                                <label className="booking-form__label" htmlFor="event-name">Название</label>
                                <input
                                    id="event-name"
                                    type="text"
                                    value={eventForm.name}
                                    onChange={(e) => setEventForm((prev) => ({ ...prev, name: e.target.value }))}
                                    required
                                />

                                <label className="booking-form__label" htmlFor="event-description">Описание</label>
                                <input
                                    id="event-description"
                                    type="text"
                                    value={eventForm.description}
                                    onChange={(e) => setEventForm((prev) => ({ ...prev, description: e.target.value }))}
                                    required
                                />

                                <label className="booking-form__label" htmlFor="event-start">Начало</label>
                                <input
                                    id="event-start"
                                    type="datetime-local"
                                    value={eventForm.startTime}
                                    onChange={(e) => setEventForm((prev) => ({ ...prev, startTime: e.target.value }))}
                                    required
                                />

                                <label className="booking-form__label" htmlFor="event-end">Окончание</label>
                                <input
                                    id="event-end"
                                    type="datetime-local"
                                    value={eventForm.endTime}
                                    onChange={(e) => setEventForm((prev) => ({ ...prev, endTime: e.target.value }))}
                                    required
                                />

                                <label className="booking-form__label" htmlFor="event-hub-id">ID хаба</label>
                                <input
                                    id="event-hub-id"
                                    type="number"
                                    min={1}
                                    value={eventForm.hubId}
                                    onChange={(e) => setEventForm((prev) => ({ ...prev, hubId: e.target.value }))}
                                    required
                                />

                                {isAdmin ? (
                                    <>
                                        <label className="booking-form__label" htmlFor="event-mentor-id">ID ментора</label>
                                        <input
                                            id="event-mentor-id"
                                            type="number"
                                            min={1}
                                            value={eventForm.mentorId}
                                            onChange={(e) => setEventForm((prev) => ({ ...prev, mentorId: e.target.value }))}
                                            required
                                        />
                                    </>
                                ) : null}

                                <button type="submit" className="btn btn--primary btn--sm" disabled={isSubmittingEvent}>
                                    {isSubmittingEvent ? "Сохраняем..." : "Создать событие"}
                                </button>
                            </form>

                            {eventSubmitMessage ? <p className="booking-card__success">{eventSubmitMessage}</p> : null}
                            {eventSubmitError ? <p className="booking-form__error">{eventSubmitError}</p> : null}
                        </article>
                    </div>
                </section>
            ) : null}

            <RegistrationModal
                isOpen={isRegistrationOpen}
                onClose={() => setIsRegistrationOpen(false)}
                eventTitle={selectedEventTitle}
                registrationType="event"
            />
        </main>
    )
}

export default EventsPage
