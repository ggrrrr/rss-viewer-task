Rails.application.routes.draw do
  get "/addresses", to: "addresses#index"
  get "/addresses/new", to: "addresses#new"
  get "/addresses/:id", to: "addresses#show", as: 'address'
  post "/addresses", to: "addresses#create"
  put "/addresses/:id", to: "addresses#update"
  delete "/addresses/:id", to: "addresses#destroy"

  # Reveal health status on /up that returns 200 if the app boots with no exceptions, otherwise 500.
  # Can be used by load balancers and uptime monitors to verify that the app is live.
  get "up" => "rails/health#show", as: :rails_health_check

  # Defines the root path route ("/")
  # root "posts#index"
  # root "#index"
  root :to => 'viewer#index'
end
