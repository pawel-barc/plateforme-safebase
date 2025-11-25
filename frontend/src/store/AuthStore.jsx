import { create } from "zustand";
import { persist } from "zustand/middleware";

const useAuthStore = create(
  persist(
    (set, get) => ({
      isAuthenticated: false,
      currentUser: null,
      // Méthode de connexion : met à jour l'utilisateur et l'état d'authentification
      login: (userData) => {
        console.log("Connexion appelée avec", userData);
        set({
          isAuthenticated: true,
          currentUser: userData,
        });
      },

      // Méthode pour vérifier si l'utilisateur est réellement connecté
      checkAuthentication: () => {
        const { isAuthenticated, currentUser } = get();
        console.log("Vérfication de l'authentification :", {
          isAuthenticated,
          hasUserData: !!currentUser,
        });
        return isAuthenticated && !!currentUser;
      },
    }),
    {
      name: "auth-storage", // Nom du stockage dans local Storage
      getStorage: () => localStorage,
    }
  )
);

export default useAuthStore;
