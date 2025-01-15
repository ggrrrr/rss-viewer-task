Rails.application.routes.draw do
  get "rss_sources/index"
  get "rss_sources/delete"
  get "rss_sources/create"
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Reveal health status on /up that returns 200 if the app boots with no exceptions, otherwise 500.
  # Can be used by load balancers and uptime monitors to verify that the app is live.
  get "up" => "rails/health#show", as: :rails_health_check

  # Render dynamic PWA files from app/views/pwa/* (remember to link manifest in application.html.erb)
  # get "manifest" => "rails/pwa#manifest", as: :pwa_manifest
  # get "service-worker" => "rails/pwa#service_worker", as: :pwa_service_worker

  # Defines the root path route ("/")
  root :to => 'home#index'

  get "/rss_sources", to: "rss_sources#index"
  get "/rss_sources/new", to: "rss_sources#new"
  get "/rss_sources/:id", to: "rss_sources#show", as: 'rss_source'
  post "/rss_sources", to: "rss_sources#create"
  put "/rss_sources/:id", to: "rss_sources#update"
  delete "/rss_sources/:id", to: "rss_sources#destroy"

end
